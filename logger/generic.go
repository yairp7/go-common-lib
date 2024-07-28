package logger

import (
	"fmt"
	"io"
)

type GenericLogger struct {
	level  LogLevel
	writer io.Writer
}

func NewGenericLogger(logLevel LogLevel, writer io.Writer) *GenericLogger {
	return &GenericLogger{
		level:  logLevel,
		writer: writer,
	}
}

func (l *GenericLogger) write(msg string, args ...any) {
	l.writer.Write([]byte(fmt.Sprintf(msg, args...)))
}

func (l *GenericLogger) Debug(msg string, args ...any) {
	if l.level > DEBUG {
		return
	}

	l.write(msg, args...)
}

func (l *GenericLogger) Info(msg string, args ...any) {
	if l.level > INFO {
		return
	}

	l.write(msg, args...)
}

func (l *GenericLogger) Warn(msg string, args ...any) {
	if l.level > WARN {
		return
	}

	l.write(msg, args...)
}

func (l *GenericLogger) Error(msg string, args ...any) {
	if l.level > ERROR {
		return
	}

	l.write(msg, args...)
}
