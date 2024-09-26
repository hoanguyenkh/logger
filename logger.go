package logger

import (
	"errors"
	"sync"
)

// A global variable so that log functions can be directly accessed
var log = DefaultLogger()

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

// LoggerBackend represents the int enum for backend of logger.
// nolint:revive
type LoggerBackend int

const (
	// Debug has verbose message
	debugLvl = "debug"
	// Info is default log level
	infoLvl = "info"
	// Warn is for logging messages about possible issues
	warnLvl = "warn"
	// Error is for logging errors
	errorLvl = "error"
	// Fatal is for logging fatal messages. The system shutdowns after logging the message.
	fatalLvl = "fatal"
)

const (
	// LoggerBackendZap logging using Uber's zap backend
	LoggerBackendZap LoggerBackend = iota
	// LoggerBackendLogrus logging using logrus backend
	LoggerBackendLogrus

	//LoggerBackendPhuslu logging using phuslu backend
	LoggerBackendPhuslu
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")

	once sync.Once
)

// Logger is our contract for the logger
type Logger interface {
	Debug(msg ...interface{})
	Debugf(format string, args ...interface{})

	Info(msg ...interface{})
	Infof(format string, args ...interface{})
	Infoln(msg string)

	Warn(msg ...interface{})
	Warnf(format string, args ...interface{})

	Error(msg ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(msg ...interface{})
	Fatalf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger

	// extract instance logger.
	GetDelegate() interface{}
	SetLogLevel(level string) error
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole    bool   `mapstructure:"enable_console"`
	EnableJSONFormat bool   `mapstructure:"enable_json_format"`
	ConsoleLevel     string `mapstructure:"console_level"`
	EnableFile       bool
	FileJSONFormat   bool
	FileLevel        string
	FileLocation     string
}

// DefaultLogger creates default logger, which uses zap sugarLogger and outputs to console
func DefaultLogger() Logger {
	cfg := Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "info",
		EnableFile:       false,
		FileJSONFormat:   false,
	}
	logger, _ := newPhusluLogger(cfg)
	return logger
}

// InitLogger returns an instance of logger
func InitLogger(config Configuration, backend LoggerBackend) (Logger, error) {
	var err error
	once.Do(func() {
		log, err = NewLogger(config, backend)
	})
	return log, err
}

func NewLogger(config Configuration, backend LoggerBackend) (Logger, error) {
	switch backend {
	case LoggerBackendZap:
		return newZapLogger(config)

	case LoggerBackendLogrus:
		return newLogrusLogger(config)

	case LoggerBackendPhuslu:
		return newPhusluLogger(config)

	default:
		return nil, errInvalidLoggerInstance
	}
}
