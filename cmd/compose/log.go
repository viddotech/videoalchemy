package main

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// VideoAlchemyLogFormatter is a Loggers formatter that adds color to log levels.
type VideoAlchemyLogFormatter struct {
	logrus.TextFormatter
}

// Format formats the log entry with colors.
func (f *VideoAlchemyLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	switch entry.Level {
	case logrus.InfoLevel:
		entry.Message = color.New(color.FgBlue).Sprint(entry.Message)
	case logrus.WarnLevel:
		entry.Message = color.New(color.FgYellow).Sprint(entry.Message)
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		entry.Message = color.New(color.FgRed).Sprint(entry.Message)
	}
	return f.TextFormatter.Format(entry)
}
