package logging

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)

	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:    true,
		CallerPrettyfier: prettierFunc,
	}

	return &Logger{
		Entry: logrus.NewEntry(logger),
	}
}

func (logger *Logger) WithField(field string, value any) *Logger {
	return &Logger{
		Entry: logger.Entry.WithField(field, value),
	}
}

func prettierFunc(frame *runtime.Frame) (function, file string) {
	file = fmt.Sprintf("%s:%v", path.Base(frame.File), frame.Line)
	function = frame.Function
	return
}
