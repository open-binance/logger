package logger

// CfgLog specifies configuration of the logger
type CfgLog struct {
	MaxSize    int    `json:"max_size" yaml:"max_size"`       // unit: MB
	MaxAge     int    `json:"max_age" yaml:"max_age"`         // unit: day
	MaxBackups int    `json:"max_backups" yaml:"max_backups"` // unit: short
	Level      string `json:"level" yaml:"level"`             // log level
	Path       string `json:"path" yaml:"path"`               // path to hold log file
	Encoding   string `json:"encoding" yaml:"encoding"`       // json or console
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
