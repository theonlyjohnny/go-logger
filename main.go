package logger

import (
	"fmt"
	"io/ioutil"
	"testing"

	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

func testLog(t *testing.T, base string, args ...interface{}) {
	if t != nil {
		t.Logf(base, args...)
	}
}

//GetLogger returns a Logger instance scoped to appName
func GetLogger(config loggerConfig) Logger {
	return realInit(config, nil)
}

func getTestLogger(config loggerConfig, t *testing.T) Logger {
	return realInit(config, t)
}

func realInit(config loggerConfig, t *testing.T) Logger {
	logger := factory.getLogger(config.appName)
	baseLogger := logger.baseLogger
	baseLogger.Level = config.logLevel

	if config.logSyslog != nil {
		testLog(t, "finalPrio: %d", config.logSyslog.logPriority)
		hook, err := logrus_syslog.NewSyslogHook("udp", config.logSyslog.remoteIP, config.logSyslog.logPriority, config.appName)
		if err == nil {
			testLog(t, "adding syslog %#v to %v -- writer: %#v", hook, baseLogger, hook.Writer)
			baseLogger.Hooks.Add(hook)

		} else {
			errString := fmt.Sprintf("Failed to add syslog %s", err.Error())
			logger.Error(errString)
			if t != nil {
				t.Error(err)
			}
		}
	}

	if !config.logConsole {
		baseLogger.Out = ioutil.Discard
	}

	return logger
}
