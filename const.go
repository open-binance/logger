package logger

const (
	maxSizeDefault  int    = 256 // unit: MB
	maxAgeDefault   int    = 7   // unit: day
	maxBackups      int    = 20  // unit: short
	levelDefault    string = "info"
	pathDefault     string = "logs"
	encodingDefault string = "console"
)
