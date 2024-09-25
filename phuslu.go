package logger

import (
	"io"

	"github.com/KyberNetwork/logger/phuslu"
	logphuslu "github.com/phuslu/log"
)

type PhusluLogger struct {
	logger    *logphuslu.Logger
	context   logphuslu.Fields
	out       io.Writer
	formatter *phuslu.Formatter
}

func newPhusluLogger(config Configuration) (*PhusluLogger, error) {
	logger := &PhusluLogger{
		logger: &logphuslu.Logger{
			Caller: 2,
		},
		context:   make(logphuslu.Fields),
		formatter: phuslu.NewFormatter(),
	}
	cfg := phuslu.Config{
		LogLevel: config.ConsoleLevel,
		Style:    phuslu.StylePretty,
	}
	logger.WithConfig(&cfg)
	if config.EnableFile && config.FileLocation != "" {
		logger.useFileWrite(config.FileLocation)
	}
	return logger, nil
}

func (l *PhusluLogger) WithFields(keyValues Fields) Logger {
	return &PhusluLogger{
		logger: &logphuslu.Logger{
			Caller: 2,
		},
		context:   logphuslu.Fields(keyValues),
		formatter: phuslu.NewFormatter(),
	}
}

func (l *PhusluLogger) WithConfig(cfg *phuslu.Config) *PhusluLogger {
	// Use default config if nil.
	if cfg == nil {
		c := phuslu.DefaultConfig()
		cfg = &c
	}
	l.withTimeFormat(cfg.TimeFormat)
	l.withStyle(cfg.Style)
	l.withLogLevel(cfg.LogLevel)
	return l
}

// withTimeFormat sets the time format for the logger.
func (l *PhusluLogger) withTimeFormat(formatStr string) {
	l.logger.TimeFormat = formatStr
}

// sets the style of the logger.
func (l *PhusluLogger) withStyle(style string) {
	if style == phuslu.StylePretty {
		l.useConsoleWriter()
	} else if style == phuslu.StyleJSON {
		l.useJSONWriter()
	}
}

// SetLevel sets the log level of the logger.
func (l *PhusluLogger) withLogLevel(level string) {
	l.logger.Level = logphuslu.ParseLevel(level)
}

// useConsoleWriter sets the logger to use a console writer.
func (l *PhusluLogger) useConsoleWriter() {
	l.setWriter(&logphuslu.ConsoleWriter{
		Writer:    l.out,
		Formatter: l.formatter.Format,
	})
}

// useConsoleWriter sets the logger to use a console writer.
func (l *PhusluLogger) useFileWrite(filename string) {
	//TODO: using AsyncWriter
	//fileWriter := logphuslu.AsyncWriter{
	//	ChannelSize: 4096,
	//	Writer: &logphuslu.FileWriter{
	//		Filename:     filename,
	//		FileMode:     0600,
	//		MaxSize:      100 * 1024 * 1024, // 100 MB
	//		MaxBackups:   100,
	//		EnsureFolder: true,
	//		LocalTime:    false,
	//	},
	//}
	fileWriter := logphuslu.FileWriter{
		Filename:     filename,
		EnsureFolder: true,
		MaxSize:      100 * 1024 * 1024, // 100 MB
		MaxBackups:   100,
		LocalTime:    false,
	}
	l.setWriter(&fileWriter)
}

// useJSONWriter sets the logger to use a IOWriter wrapper.
func (l *PhusluLogger) useJSONWriter() {
	l.setWriter(logphuslu.IOWriter{Writer: l.out})
}

// setWriter sets the writer of the logger.
func (l *PhusluLogger) setWriter(writer logphuslu.Writer) {
	l.logger.Writer = writer
}

func (l *PhusluLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *PhusluLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *PhusluLogger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *PhusluLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *PhusluLogger) Infoln(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *PhusluLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *PhusluLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *PhusluLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *PhusluLogger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *PhusluLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}

func (l *PhusluLogger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

func (l *PhusluLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

func (l *PhusluLogger) Panic(msg string) {
	l.logger.Panic().Msg(msg)
}

func (l *PhusluLogger) GetDelegate() interface{} {
	return l.logger
}

func (l *PhusluLogger) SetLogLevel(logLevel string) error {
	l.logger.Level = logphuslu.ParseLevel(logLevel)
	return nil
}
