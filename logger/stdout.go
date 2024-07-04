package logger

import "fmt"

type StdoutLogger struct {
	level LogLevel
}

func NewStdoutLogger(logLevel LogLevel) *StdoutLogger {
	return &StdoutLogger{
		level: logLevel,
	}
}

func (l *StdoutLogger) Debug(msg string, args ...any) {
	if l.level > DEBUG {
		return
	}

	fmt.Println(fmt.Sprintf(msg, args...))
}

func (l *StdoutLogger) Info(msg string, args ...any) {
	if l.level > INFO {
		return
	}

	fmt.Println(fmt.Sprintf(msg, args...))
}

func (l *StdoutLogger) Warn(msg string, args ...any) {
	if l.level > WARN {
		return
	}

	fmt.Println(fmt.Sprintf(msg, args...))
}

func (l *StdoutLogger) Error(msg string, args ...any) {
	if l.level > ERROR {
		return
	}

	fmt.Println(fmt.Sprintf(msg, args...))
}
