package rlog

import (
	"fmt"
	"time"
)

type Logger struct {
	Name string
	Mode string
}

func NewLogger(mode string, name string) *Logger {
	return &Logger{Mode: mode, Name: name}
}

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Logger) Info(message string) {
	if l.Mode == "dev" {
		fmt.Printf("%s {%s} [INFO] %s\n", getTime(), l.Name, message)
	}
}

func (l *Logger) Fatal(message string) {
	str := fmt.Sprintf("%s {%s} [FATAL] %s\n", getTime(), l.Name, message)
	panic(str)
}

func (l *Logger) Debug(message string) {
	if l.Mode == "dev" {
		fmt.Printf("%s {%s} [DEBUG] %s\n", getTime(), l.Name, message)
	}
}
