package golog

import (
	"log"
	"log/syslog"
	"net/url"
	"os"
)

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
	if to == "" {
		Infoln("Log to STDERR")
		return
	}

	url, err := url.Parse(to)
	if err != nil {
		log.Fatalln("Error happened when parse logging URL ", to, ":", err)
	}

	// File URL must contain only `url.Path`. Syslog location must contain only `url.Host`
	if (url.Host == "" && url.Path == "") || (url.Host != "" && url.Path != "") {
		log.Fatalln("Logging location is wrong:", to)
	}

	switch url.Scheme {
	case "udp", "tcp": // FIXME: must it be syslog://address.syslog:port ?
		writer, err := syslog.Dial(url.Scheme, url.Host, syslog.LOG_INFO, prefix)
		if err != nil {
			log.Fatalln("Unable to connect to ", to, " : ", err)
		}
		log.SetOutput(writer)
	case "file":
		fwriter, err := os.OpenFile(url.Path, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Cannot open file '", url.Path, "':", err)
		}
		log.SetOutput(fwriter)
	default:
		log.Fatalln("Unknown logging location protocol:", url.Scheme)
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

func (l *Logger) Fatalln(a ...interface{}) {
	a = append([]interface{}{"FATAL --"}, a...)
	log.Println(a...)
	os.Exit(1)
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

func Fatalln(a ...interface{}) {
	DefaultLogger.Fatalln(a...)
}
