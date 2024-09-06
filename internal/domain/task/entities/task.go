package entities

import (
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
)

type TaskID string
type TaskStatus string

const (
	STARTED TaskStatus = "STARTED"
	RUNNING TaskStatus = "RUNNING"
	DONE    TaskStatus = "DONE"
	FAILED  TaskStatus = "FAILED"
)

type Task struct {
	ID                   TaskID
	Instruction          schema.Instruction
	FFMPEGCommand        []string
	Status               TaskStatus
	GeneratedCommandPath string
	GeneratedLogPath     string
}
