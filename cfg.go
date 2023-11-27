package logger

import "fmt"

// CfgLog specifies configuration of the logger
type CfgLog struct {
	MaxSize    int    `json:"max_size" yaml:"max_size"`       // unit: MB
	MaxAge     int    `json:"max_age" yaml:"max_age"`         // unit: day
	MaxBackups int    `json:"max_backups" yaml:"max_backups"` // unit: short
	Level      string `json:"level" yaml:"level"`             // log level
	Path       string `json:"path" yaml:"path"`               // path to hold log file
	Encoding   string `json:"encoding" yaml:"encoding"`       // json or console
}

// FileConfig specifies detail file config
type FileConfig struct {
	Enable     bool   `json:"enable" yaml:"enable"`           // it will not write message to file in case of false
	LocalTime  bool   `json:"local_time" yaml:"local_time"`   // it determines if the time used for formatting the timestamps in backup files is the computer's local time. The default is to use UTC time
	Compress   bool   `json:"compress" yaml:"compress"`       // it determines if the rotated log files should be compressed using gzip. The default is not to perform compression
	MaxSize    int    `json:"max_size" yaml:"max_size"`       // the maximum size in megabytes of the log file before it gets, it defaults to 100 megabytes
	MaxAge     int    `json:"max_age" yaml:"max_age"`         // the maximum number of days to retain old log files based on the timestamp encoded in their filename
	MaxBackups int    `json:"max_backups" yaml:"max_backups"` // the maximum number of old log files to retain
	Level      string `json:"level" yaml:"level"`             // the log level
	Filename   string `json:"filename" yaml:"filename"`       // the file to write logs to
	TimeKey    string `json:"time_key" yaml:"time_key"`       // set "" is ok
	CallerKey  string `json:"caller_key" yaml:"caller_key"`   // set "" is ok
	LevelKey   string `json:"level_key" yaml:"level_key"`     // set "" is ok
}

// NewDefaultCfg creates a new logger config by default
func NewDefaultCfg() CfgLog {
	return CfgLog{
		MaxSize:    maxSizeDefault,
		MaxAge:     maxAgeDefault,
		MaxBackups: maxBackups,
		Level:      levelDefault,
		Path:       pathDefault,
		Encoding:   encodingDefault,
	}
}

func (cl *CfgLog) SetLevel(level string) error {
	if level != gDebug && level != gInfo && level != gWarn && level != gError {
		return fmt.Errorf("unknown level: %s, only support %s, %s, %s, %s", level, gDebug, gInfo, gWarn, gError)
	}

	cl.Level = level
	return nil
}

func NewCfg(maxSize, maxAge, maxBackups int, level, path, encoding string) CfgLog {
	return CfgLog{
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		Level:      level,
		Path:       path,
		Encoding:   encoding,
	}
}
