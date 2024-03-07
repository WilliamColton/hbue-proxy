package log

import (
	"os"
	"strings"
)

type Log struct {
	LogSrc string
}

func NewLog(logSrc string) *Log {
	return &Log{LogSrc: logSrc}
}

func (l *Log) RecordError(err error) {
	var f *os.File
	f, _ = os.OpenFile(l.LogSrc, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	f.Seek(0, 2)

	logData := err.Error()
	if strings.HasSuffix(logData, "\n") {
		f.WriteString(logData)
	} else {
		f.WriteString(logData + "\n")
	}
	f.Sync()
	f.Close()
}
