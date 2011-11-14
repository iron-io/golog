package golog

import (
	"log"
)

const (
	Debug = iota
	Info
	Warn
	Error
)

type Logger struct {
	Level int
}

func NewLogger() *Logger {
	l := new(Logger)
	l.Level = Info
	return l
}

func (l *Logger) Debugln(a ...interface{}) {
	if l.Level <= Debug {
		log.Println(a...)
	}
}
func (l *Logger) Infoln(a ...interface{}) {
	if l.Level <= Info {
		log.Println(a...)
	}
}
