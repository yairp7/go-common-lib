package logger

import "os"

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

func NewMixedLogger(opts ...MixedLoggerOption) *MixedLogger {
	logger := &MixedLogger{
		level: DEBUG,
	}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func (l *MixedLogger) Debug(msg string, args ...any) {
	if l.level > DEBUG {
		return
	}

	for _, logger := range l.loggerImpl {
		logger.Debug(msg, args...)
	}
}

func (l *MixedLogger) Info(msg string, args ...any) {
	if l.level > INFO {
		return
	}

	for _, logger := range l.loggerImpl {
		logger.Info(msg, args...)
	}
}

func (l *MixedLogger) Warn(msg string, args ...any) {
	if l.level > WARN {
		return
	}

	for _, logger := range l.loggerImpl {
		logger.Warn(msg, args...)
	}
}

func (l *MixedLogger) Error(msg string, args ...any) {
	if l.level > ERROR {
		return
	}

	for _, logger := range l.loggerImpl {
		logger.Error(msg, args...)
	}
}

func NewStdoutLogger(logLevel LogLevel) *GenericLogger {
	return &GenericLogger{
		level:  logLevel,
		writer: os.Stdout,
	}
}
