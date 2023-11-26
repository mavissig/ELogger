package logger

import "log"

type Logger struct {
}

func (l *Logger) Mock() {
	log.Println("Mock call")
}

func New() *Logger {
	return &Logger{}
}
