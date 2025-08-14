package logger

import (
	"log"
	"time"
)

type LogEntry struct {
	Time    time.Time
	Action  string
	Message string
}

type Logger struct {
	logChan chan LogEntry
}

func NewLogger() *Logger {
	logger := &Logger{
		logChan: make(chan LogEntry, 10),
	}
	go logger.processLogs()
	return logger
}

func (l *Logger) Log(action, message string) {
	l.logChan <- LogEntry{
		Time:    time.Now(),
		Action:  action,
		Message: message,
	}
}

func (l *Logger) processLogs() {
	for entry := range l.logChan {
		log.Printf("[%s] %s: %s\n", entry.Time.Format(time.RFC3339), entry.Action, entry.Message)
	}
}

func (l *Logger) Close() {
	close(l.logChan)
}
