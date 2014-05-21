package golog

import (
	"fmt"
	"log"
	"log/syslog"
	"net/url"
	"os"
)

const (
	Debug = iota
	Info
	Warn
	Error
	Fatal
)

func SetLogLevel(level string) {
	switch level {
	case "debug":
		DefaultLogger.Level = Debug
	case "warn":
		DefaultLogger.Level = Warn
	case "error":
		DefaultLogger.Level = Error
	case "fatal":
		DefaultLogger.Level = Fatal
	default:
		DefaultLogger.Level = Info
	}
}

func levelName(level int) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"

	}

	return "???"
}

func SetLogLocation(to, prefix string) {
	if to == "" {
		Infoln("Log to STDERR")
		return
	}

	pUrl, err := url.Parse(to)
	if err != nil {
		log.Fatalln("Error happened when parse logging URL ", to, ":", err)
	}

	if pUrl.Host == "" && pUrl.Path == "" {
		log.Println("No scheme on logging url, adding udp://")
		// this happens when no scheme like udp:// is present
		to = fmt.Sprintf("udp://%v", to)
		pUrl, err = url.Parse(to)
		if err != nil {
			log.Fatalln("Error happened when parse logging URL ", to, ":", err)
		}
	}

	// File URL must contain only `url.Path`. Syslog location must contain only `url.Host`
	if (pUrl.Host == "" && pUrl.Path == "") || (pUrl.Host != "" && pUrl.Path != "") {
		log.Fatalln("Invalid logging location:", to)
	}

	switch pUrl.Scheme {
	case "udp", "tcp": // FIXME: must it be syslog://address.syslog:port ?
		writer, err := syslog.Dial(pUrl.Scheme, pUrl.Host, syslog.LOG_INFO, prefix)
		if err != nil {
			log.Fatalln("Unable to connect to ", to, " : ", err)
		}
		log.SetOutput(writer)
	case "file":
		fwriter, err := os.OpenFile(pUrl.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Cannot open file '", pUrl.Path, "':", err)
		}
		log.SetOutput(fwriter)
	default:
		log.Fatalln("Unknown logging location protocol:", pUrl.Scheme)
	}
}

var DefaultLogger = NewLogger()

type Logger struct {
	Level int
}

func NewLogger() *Logger {
	l := new(Logger)
	l.Level = Info
	return l
}

func (l *Logger) Logln(level int, a ...interface{}) {
	if l.Level <= level {
		a = append([]interface{}{fmt.Sprintf("%v --", levelName(level))}, a...)
		log.Println(a...)
	}
}

func (l *Logger) Logf(level int, format string, a ...interface{}) {
	if l.Level <= level {
		// probably a better way to do this
		log.Printf(fmt.Sprintf("%v --", levelName(level))+format, a...)
	}
}

func (l *Logger) Debugln(a ...interface{}) {
	l.Logln(Debug, a)
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Logf(Debug, format, a)
}

func (l *Logger) Infoln(a ...interface{}) {
	l.Logln(Info, a)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Logf(Info, format, a)
}

func (l *Logger) Warnln(a ...interface{}) {
	l.Logln(Warn, a)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Logf(Warn, format, a)
}

func (l *Logger) Errorln(a ...interface{}) {
	l.Logln(Error, a)
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Logf(Error, format, a)
}

func (l *Logger) Fatalln(a ...interface{}) {
	l.Logln(Fatal, a)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Logf(Fatal, format, a)
	os.Exit(1)
}

func Logln(level int, a ...interface{}) {
	DefaultLogger.Logln(level, a)
}

func Logf(level int, format string, a ...interface{}) {
	DefaultLogger.Logf(level, format, a)
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

func Fatalf(format string, a ...interface{}) {
	DefaultLogger.Fatalf(format, a...)
}
