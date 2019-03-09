package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type loggerFactory struct {
	loggers map[string]Logger
}

var factory = &loggerFactory{
	map[string]Logger{},
}

func (f *loggerFactory) getLogger(appName string) Logger {
	if logger, ok := f.loggers[appName]; ok {
		return logger
	}

	logger := Logger{
		baseLogger: &logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(CustomFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
		AppName: appName,
	}

	f.loggers[appName] = logger

	return f.getLogger(appName)
}
