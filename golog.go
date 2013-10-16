package golog

import (
	"log"
	"log/syslog"
)

type LoggingConfig struct {
	To     string `json:"to"`
	Level  string `json:"level"`
	Prefix string `json:"prefix"`
}

func SetLogLevel(level string) {
	switch level {
	case "debug":
		DefaultLogger.Level = Debug
	case "warn":
		DefaultLogger.Level = Warn
	case "error":
		DefaultLogger.Level = Error
	default:
		DefaultLogger.Level = Info
	}
}

func SetLogLocation(to, prefix string) {
	switch to {
	case "":
	default:
		writer, err := syslog.Dial("udp", to, syslog.LOG_INFO, prefix)
		if err != nil {
			log.Fatalln("unable to connect to ", to, " : ", err)
		}
		log.SetOutput(writer)
	}
}

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
		a = append([]interface{}{"DEBUG --"}, a...)
		log.Println(a...)
	}
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.Level <= Debug {
		log.Printf("DEBUG -- "+format, a...)
	}
}

func (l *Logger) Infoln(a ...interface{}) {
	if l.Level <= Info {
		a = append([]interface{}{"INFO --"}, a...)
		log.Println(a...)
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	if l.Level <= Info {
		log.Printf("INFO -- "+format, a...)
	}
}

func (l *Logger) Warnln(a ...interface{}) {
	if l.Level <= Warn {
		a = append([]interface{}{"WARN --"}, a...)
		log.Println(a...)
	}
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	if l.Level <= Warn {
		log.Printf("WARN -- "+format, a...)
	}
}

func (l *Logger) Errorln(a ...interface{}) {
	if l.Level <= Error {
		a = append([]interface{}{"ERROR --"}, a...)
		log.Println(a...)
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	if l.Level <= Error {
		log.Printf("ERROR -- "+format, a...)
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
