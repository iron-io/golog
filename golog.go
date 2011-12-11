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

var DefaultLogger = NewLogger()

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
		// doesn't work: a = append([]interface{"DEBUG --"}, a...)
		log.Println(a...)
	}
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.Level <= Debug {
		log.Printf(format, a...)
	}
}

func (l *Logger) Infoln(a ...interface{}) {
	if l.Level <= Info {
		log.Println(a...)
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	if l.Level <= Info {
		log.Printf(format, a...)
	}
}

func (l *Logger) Warnln(a ...interface{}) {
	if l.Level <= Warn {
		log.Println(a...)
	}
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	if l.Level <= Warn {
		log.Printf(format, a...)
	}
}

func (l *Logger) Errorln(a ...interface{}) {
	if l.Level <= Error {
		log.Println(a...)
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	if l.Level <= Error {
		log.Printf(format, a...)
	}
}

func Debugln(a ...interface{}) {
	DefaultLogger.Debugln(a...)
}

func Debugf(format string, a ...interface{}) {
	DefaultLogger.Debugf(format, a...)
}

func Infoln(a ...interface{}) {
	DefaultLogger.Infoln(a...)
}

func Infof(format string, a ...interface{}) {
	DefaultLogger.Infof(format, a...)
}

func Warnln(a ...interface{}) {
	DefaultLogger.Warnln(a...)
}

func Warnf(format string, a ...interface{}) {
	DefaultLogger.Warnf(format, a...)
}

func Errorln(a ...interface{}) {
	DefaultLogger.Errorln(a...)
}

func Errorf(format string, a ...interface{}) {
	DefaultLogger.Errorf(format, a...)
}
