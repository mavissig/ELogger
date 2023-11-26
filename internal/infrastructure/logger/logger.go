package logger

import (
	"elogger/internal/domain"
	"encoding/json"
	"fmt"
	"log"
)

type Logger struct {
}

func (l *Logger) Mock() {
	log.Println("Mock call")
}

func (l *Logger) LogIP(l_ domain.LogIP__) {
	jsonByte, err := json.MarshalIndent(l_, "", "    ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(jsonByte))
}

func New() *Logger {
	return &Logger{}
}
