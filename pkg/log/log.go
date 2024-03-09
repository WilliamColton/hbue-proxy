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

func (l *Log) RecordError(err error) error {
	f, err := os.OpenFile(l.LogSrc, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := f.Seek(0, 2); err != nil {
		return err
	}

	logData := err.Error()
	if strings.HasSuffix(logData, "\n") {
		if _, err := f.WriteString(logData); err != nil {
			return err
		}
	} else {
		if _, err := f.WriteString(logData + "\n"); err != nil {
			return err
		}
	}
	if err := f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
