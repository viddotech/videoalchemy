package entities

import (
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"time"
)

type TaskID string

type Task struct {
	ID            TaskID
	Instruction   schema.Instruction
	FFMPEGCommand *string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
