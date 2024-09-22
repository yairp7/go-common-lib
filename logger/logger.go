package logger

import (
	"fmt"
	"os"
)

const (
	DEBUG LogLevel = iota
	INFO  LogLevel = iota
	WARN  LogLevel = iota
	ERROR LogLevel = iota
)

type LogLevel int

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type MixedLoggerOption func(*MixedLogger)

type MixedLogger struct {
	loggerImpl []Logger
	level      LogLevel
	suffix     string
}

func WithLoggerImpl(loggerImpl Logger) MixedLoggerOption {
	return func(ml *MixedLogger) {
		ml.loggerImpl = append(ml.loggerImpl, loggerImpl)
	}
}

func WithLoggerLevel(level LogLevel) MixedLoggerOption {
	return func(ml *MixedLogger) {
		ml.level = level
	}
}

func WithLogSuffix(suffix string) MixedLoggerOption {
	return func(ml *MixedLogger) {
		ml.suffix = suffix
	}
}

func NewMixedLogger(opts ...MixedLoggerOption) *MixedLogger {
	logger := &MixedLogger{
		level: DEBUG,
	}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func (l *MixedLogger) buildMessage(msg string) string {
	return fmt.Sprintf("%s%s", msg, l.suffix)
}

func (l *MixedLogger) Debug(msg string, args ...any) {
	if l.level > DEBUG {
		return
	}

	fMsg := l.buildMessage(msg)

	for _, logger := range l.loggerImpl {
		logger.Debug(fMsg, args...)
	}
}

func (l *MixedLogger) Info(msg string, args ...any) {
	if l.level > INFO {
		return
	}

	fMsg := l.buildMessage(msg)

	for _, logger := range l.loggerImpl {
		logger.Info(fMsg, args...)
	}
}

func (l *MixedLogger) Warn(msg string, args ...any) {
	if l.level > WARN {
		return
	}

	fMsg := l.buildMessage(msg)

	for _, logger := range l.loggerImpl {
		logger.Warn(fMsg, args...)
	}
}

func (l *MixedLogger) Error(msg string, args ...any) {
	if l.level > ERROR {
		return
	}

	fMsg := l.buildMessage(msg)

	for _, logger := range l.loggerImpl {
		logger.Error(fMsg, args...)
	}
}

func NewStdoutLogger(logLevel LogLevel) *GenericLogger {
	return &GenericLogger{
		level:  logLevel,
		writer: os.Stdout,
	}
}
