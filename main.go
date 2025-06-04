package main

import (
	"bytes"
	"fmt"
	"log"
)

type LogLevel int

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	logger   *log.Logger
	logLevel LogLevel
}

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func NewLogExtended(buf *bytes.Buffer) LogExtended {
	logger := log.New(buf, "logger :", log.Lshortfile)
	return LogExtended{logger: logger, logLevel: LogLevelWarning}
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERROR ", msg)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	//fmt.Printf("srcLevel: %d, set level: %v\n", srcLogLvl, l.logLevel)
	if srcLogLvl > l.logLevel {
		return
	}
	l.logger.Println(prefix + msg)
}

func main() {
	var buffer bytes.Buffer
	loggerOut := NewLogExtended(&buffer)
	loggerOut.SetLogLevel(LogLevelWarning)
	loggerOut.Errorln("this is an error")
	loggerOut.Warnln("this is a warning")
	loggerOut.Infoln("you will not see this")
	fmt.Println(&buffer)
}
