package logger

import "fmt"

type StdoutLogger struct{}

func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{}
}

func (l *StdoutLogger) Debug(msg string, args ...any) {
	fmt.Println(fmt.Sprintf(msg, args...))
}

func (l *StdoutLogger) Info(msg string, args ...any) {
	fmt.Println(fmt.Sprintf(msg, args...))
}

func (l *StdoutLogger) Warn(msg string, args ...any) {
	fmt.Println(fmt.Sprintf(msg, args...))
}

func (l *StdoutLogger) Error(msg string, args ...any) {
	fmt.Println(fmt.Sprintf(msg, args...))
}
