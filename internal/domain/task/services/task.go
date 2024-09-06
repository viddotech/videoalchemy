package services

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/viddotech/videoalchemy/internal/domain/task/entities"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"github.com/viddotech/videoalchemy/internal/infrastructure/pretty"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type TaskService struct {
	Tasks []*entities.Task
}

func (s *TaskService) CreateTasks(allInstructions []schema.Instruction, GeneratePath string) error {

	for _, instruction := range allInstructions {
		ffmpegCommand, err := ffmpeg.GenerateFFMPEGCommand(instruction, allInstructions, GeneratePath)
		if err != nil {
			log.Fatal("Error generate Ffmpeg command : ", err)
		}
		task := &entities.Task{
			ID:            entities.TaskID(uuid.New().String()),
			Instruction:   instruction,
			FFMPEGCommand: ffmpegCommand,
		}
		s.Tasks = append(s.Tasks, task)

	}
	return nil
}

func (s *TaskService) RunTasks(generatePath string) bool {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(len(s.Tasks))

	for _, task := range s.Tasks {
		go s.Run(task, generatePath, ctx, cond, &wg, cancel)
	}

	wg.Wait() // Wait for all tasks to complete
	for _, task := range s.Tasks {
		if task.Status == entities.FAILED {
			return false
		}
	}

	return true
}

func (s *TaskService) Run(task *entities.Task, generatePath string, ctx context.Context, cond *sync.Cond, wg *sync.WaitGroup, cancel context.CancelFunc) {
	task.Status = entities.STARTED
	s.Notify(task)

	defer wg.Done()

	cond.L.Lock()

	for !s.AllRelatedTasksDone(task) {
		select {
		case <-ctx.Done():
			if s.RelatedTasksFailed(task) {
				return
			}
		default:
			cond.Wait()
		}

	}

	cmd := exec.Command(task.FFMPEGCommand[0], task.FFMPEGCommand[1:]...)

	// Print FFMPEG Generated Command

	defer cond.L.Unlock()

	// Get a pipe to the command's standard output
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("Error creating StdoutPipe:", err)
		task.Status = entities.FAILED
		s.Notify(task)
		cancel()
		cond.Broadcast()
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Errorf("Error creating StderrPipe: %v\n", err)
		task.Status = entities.FAILED
		s.Notify(task)
		cancel()
		cond.Broadcast()
		return
	}

	logFile := s.CreateLogFile(task, generatePath, "command.log")
	task.GeneratedLogPath = logFile.Name()

	defer logFile.Close()
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	commandFile := s.CreateLogFile(task, generatePath, "command.txt")
	task.GeneratedCommandPath = commandFile.Name()
	_, err = commandFile.WriteString(fmt.Sprintf("%s\n", strings.Join(task.FFMPEGCommand, " ")))
	if err != nil {
		log.Fatal("Error writing command to file:", err)
	}
	defer commandFile.Close()

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Error("Error starting command:", err)
		task.Status = entities.FAILED
		s.Notify(task)
		cancel()
		cond.Broadcast()
		return
	}

	task.Status = entities.RUNNING
	s.Notify(task)
	// Use another goroutine to read from the stdoutPipe and print the output
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// Use another goroutine to read from the StderrPipe and print the output
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Errorf("Error command %s : %s", task.Instruction.Name, err)
		task.Status = entities.FAILED
		s.Notify(task)
		cancel()
		cond.Broadcast()
		return
	}

	task.Status = entities.DONE
	s.Notify(task)
	cond.Broadcast() // Wake up all waiting tasks
}

func (s *TaskService) AllRelatedTasksDone(task *entities.Task) bool {
	for _, name := range task.Instruction.RunAfter {
		for _, t := range s.Tasks {
			if t.Instruction.Name == name && t.Status != entities.DONE {
				return false
			}
		}
	}
	return true
}

func (s *TaskService) RelatedTasksFailed(task *entities.Task) bool {
	for _, name := range task.Instruction.RunAfter {
		for _, t := range s.Tasks {
			if t.Instruction.Name == name && t.Status == entities.FAILED {
				return true
			}
		}
	}
	return false
}

func (s *TaskService) CreateLogFile(task *entities.Task, generatePath string, fileName string) *os.File {
	absPath, err := filepath.Abs(generatePath)
	basePath := fmt.Sprintf("%s/%s", absPath, task.Instruction.Name)
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		err := os.MkdirAll(basePath, 0750)
		if err != nil {
			log.Fatal("Error creating directory:", err)
		}
	}
	if err != nil {
		log.Fatal("Error getting absolute path:", err)
	}
	file, err := os.Create(fmt.Sprintf("%s/%s", basePath, fileName))
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	return file
}

func (s *TaskService) Notify(task *entities.Task) {

	if task.Status == entities.DONE {
		pretty.NotifySuccessText("Yay! Task %s is %s 🎉\n\nGenerated Files:\n- Log file: %s\n- Command file: %s", task.Instruction.Name, string(task.Status), task.GeneratedLogPath, task.GeneratedCommandPath)
	} else if task.Status == entities.FAILED {
		pretty.NotifyDangerousText("Oops! Task %s is %s 😢\n\nGenerated Files:\n- Log file: %s\n- Command file: %s", task.Instruction.Name, string(task.Status), task.GeneratedLogPath, task.GeneratedCommandPath)
	} else if task.Status == entities.RUNNING {
		pretty.NotifyNormalText("Task %s is %s 🚀\n\nGenerated Files:\n- Log file: %s\n- Command file: %s", task.Instruction.Name, string(task.Status), task.GeneratedLogPath, task.GeneratedCommandPath)
	} else {
		pretty.NotifyNormalText("Task %s is %s", task.Instruction.Name, string(task.Status))
	}
}
