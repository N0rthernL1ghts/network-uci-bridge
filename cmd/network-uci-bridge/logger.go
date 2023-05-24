package main

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	return &Logger{logger: logger}
}

func (l *Logger) Error(msg string) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Println(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatalln(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.SetPrefix("INFO: ")
	l.logger.Println(msg)
}
