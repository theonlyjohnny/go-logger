package logger

import (
	"log"
	"log/syslog"
	"testing"

	"github.com/sirupsen/logrus"
)

const (
	remote string = "localhost:514"
)

func assert(t *testing.T, prop string, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf("%s mismatch. Expected: %#v, actual: %#v", prop, expected, actual)
	}
}

func TestGetAllLogLevels(t *testing.T) {
	appName := "TestGetValidConfig"
	levels := logrus.AllLevels
	for _, level := range levels {
		levelNameBytes, err := level.MarshalText()
		if err != nil {
			continue
		}

		levelName := string(levelNameBytes)

		config, err := NewConfig(appName, levelName, true, nil)
		if err != nil {
			log.Fatalf("%s, %s", level, err.Error())
		}
		assert(t, appName+"["+levelName+"]", level, config.logLevel)
	}
}

func TestGetValidConsoleConfig(t *testing.T) {
	appName := "TestGetValidConfig"
	config, err := NewConfig(appName, "debug", true, nil)
	t.Logf("TestGetValidConfig config: %#v", config)
	if err != nil {
		t.Fatal(err)
	}

	var logSyslog *syslogConfig
	assert(t, "config.appName", appName, config.appName)
	assert(t, "config.logLevel", logrus.DebugLevel, config.logLevel)
	assert(t, "config.logConsole", true, config.logConsole)
	assert(t, "config.logSyslog", logSyslog, config.logSyslog)
}

func TestGetInvalidConsoleConfig(t *testing.T) {
	appName := t.Name()
	config, err := NewConfig(appName, "debug", false, nil)
	t.Logf("%s config: %#v", appName, config)
	if err == nil {
		t.Fatal("err shouldn't be nil")
	}
}

func TestGetValidSyslogConfig(t *testing.T) {
	appName := t.Name()
	syslogConfig := SyslogConfig{
		remote,
		"debug",
	}
	config, err := NewConfig(appName, "debug", false, &syslogConfig)
	t.Logf("%s config: %#v", appName, config)
	if err != nil {
		t.Fatal(err)
	}
	assert(t, "config.appName", appName, config.appName)
	assert(t, "config.logLevel", logrus.DebugLevel, config.logLevel)
	assert(t, "config.logConsole", false, config.logConsole)
	assert(t, "config.logSyslog.remoteIP", syslogConfig.RemoteIP, config.logSyslog.remoteIP)
	assert(t, "config.logSyslog.logPriority", syslog.LOG_DEBUG, config.logSyslog.logPriority)
}

func TestGetInvalidSyslogConfig(t *testing.T) {
	appName := t.Name()
	syslogConfig := SyslogConfig{}
	config, err := NewConfig(appName, "debug", false, &syslogConfig)
	t.Logf("%s config: %#v", appName, config)
	if err == nil {
		t.Fatal("err shouldn't be nil")
	}
}

func TestGetLogger(t *testing.T) {
	appName := "TestGetValidConfig"
	config, err := NewConfig(appName, "debug", true, nil)
	t.Logf("TestGetValidConfig config: %#v", config)
	if err != nil {
		t.Fatal(err)
	}

	var logSyslog *syslogConfig
	assert(t, "config.appName", appName, config.appName)
	assert(t, "config.logLevel", logrus.DebugLevel, config.logLevel)
	assert(t, "config.logConsole", true, config.logConsole)
	assert(t, "config.logSyslog", logSyslog, config.logSyslog)
	logger := GetLogger(*config)
	logger.Info("info")
	logger.Infof("infof %s %v %#v %d %b", "string", "vString", "complexVString", 5, 5)
	logger.Warn("warn")
	logger.Warnf("warnf %s %v %#v %d %b", "string", "vString", "complexVString", 5, 5)
	logger.Error("error")
	logger.Errorf("errorf %s %v %#v %d %b", "string", "vString", "complexVString", 5, 5)
	logger.Debug("debug")
	logger.Debugf("debugf %s %v %#v %d %b", "string", "vString", "complexVString", 5, 5)
}

// func TestAllLogs(t *testing.T) {
// logger := getTestLogger("bebo-go-commons-TestAllLogs2", remote, "debug", "debug", true, t)

// logger.Debug("debug")
// logger.Info("info")
// logger.Warn("warn")
// logger.Error("error")
// }

// func TestAllLogFormat(t *testing.T) {

// logger := getTestLogger("bebo-go-commons-TestAllLogFormat", remote, "debug", "debug", true, t)

// rn := time.Now().String()

// logger.Debugf("debug @ %s", rn)
// logger.Infof("info @ %s", rn)
// logger.Warnf("warn @ %s", rn)
// logger.Error("error @ %s", rn)

// }

// func TestTwoLoggers(t *testing.T) {
// loggerOne := getTestLogger("bebo-go-commons-TestTwoLoggers-1", remote, "debug", "debug", true, t)
// loggerTwo := getTestLogger("bebo-go-commons-TestTwoLoggers-2", remote, "debug", "debug", true, t)

// loggerOne.Infof("loggerOne")
// loggerTwo.Infof("loggerTwo")
// }

// func TestNoSyslog(t *testing.T) {
// logger := getTestLogger("bebo-go-commons-TestNoSyslog", remote, "debug", "", true, t)
// logger.Info("not in syslog :)")
// }

// func TestChooseRsyslogServer(t *testing.T) {
// logger := getTestLogger("bebo-go-commons-TestChooseRsyslogServer", "", "debug", "debug", true, t)
// logger.Info("Where am I?")
// }
