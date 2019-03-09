package logger

import (
	"errors"
	"log/syslog"

	"github.com/sirupsen/logrus"
)

var levelToPriority = map[string]syslog.Priority{
	"emerg":   syslog.LOG_EMERG, /* system is unusable */
	"alert":   syslog.LOG_ALERT, /* action must be taken immediately */
	"crit":    syslog.LOG_CRIT,  /* critical conditions */
	"err":     syslog.LOG_ERR,   /* error conditions */
	"error":   syslog.LOG_ERR,
	"warning": syslog.LOG_WARNING, /* warning conditions */
	"warn":    syslog.LOG_WARNING, /* warning conditions */
	"notice":  syslog.LOG_NOTICE,  /* normal but significant condition */
	"info":    syslog.LOG_INFO,    /* informational */
	"debug":   syslog.LOG_DEBUG,   /* debug-level messages */
}

type loggerConfig struct {
	appName    string
	logLevel   logrus.Level
	logConsole bool
	logSyslog  *syslogConfig
}

//SyslogConfig is a sub-config for syslog configuration
type SyslogConfig struct {
	RemoteIP    string
	LogPriority string
}

type syslogConfig struct {
	remoteIP    string
	logPriority syslog.Priority
}

//NewConfig creates a Config instance which determines how/where the logger displays messages
func NewConfig(appName, logLevel string, logConsole bool, logSyslog *SyslogConfig) (*loggerConfig, error) {
	if appName == "" {
		return nil, errors.New("appName is required")
	}
	if logLevel == "" {
		return nil, errors.New("logLevel is required")
	}

	logLvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, errors.New("Invalid logLevel")
	}
	var realLogSyslog *syslogConfig

	if logSyslog != nil {
		if logSyslog.RemoteIP == "" {
			return nil, errors.New("If logSyslog is specified, logSyslog.RemoteIP is required")
		}
		if logSyslog.LogPriority == "" {
			return nil, errors.New("If logSyslog is specified, logSyslog.LogPriority is required")
		}
		logPrio, ok := levelToPriority[logSyslog.LogPriority]
		if !ok {
			return nil, errors.New("Invalid logPriority")
		}
		realLogSyslog = &syslogConfig{
			remoteIP:    logSyslog.RemoteIP,
			logPriority: logPrio,
		}
	}

	if logSyslog == nil && !logConsole {
		return nil, errors.New("Must have at least either logSyslog or logConsole enabled")
	}
	c := loggerConfig{
		appName,
		logLvl,
		logConsole,
		realLogSyslog,
	}
	return &c, nil
}
