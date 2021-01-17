package hooks

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Hook struct will implement the logrus
// This will buffer logs and write logs at appropriate severity
type Hook struct {
	Writer  io.Writer
	Linebuf [][]byte
}

// Fire function will trigger whenever logged
func (hook *Hook) Fire(entry *logrus.Entry) error {
	bytes, err := entry.Bytes()
	if err != nil {
		return err
	}
	hook.Linebuf = append(hook.Linebuf, bytes)

	if entry.Level == logrus.ErrorLevel {
		var writeErr error
		for _, line := range hook.Linebuf {
			_, writeErr = hook.Writer.Write(line)
		}
		// clear the buffer
		hook.Linebuf = nil
		return writeErr
	}
	return nil
}

// Levels set a trigger to fire at appropriate log levels
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
