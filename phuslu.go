package logger

import (
	"io"

	"github.com/KyberNetwork/logger/phuslu"
	logphuslu "github.com/phuslu/log"
)

type phusluLogger struct {
	logger    *logphuslu.Logger
	context   logphuslu.Fields
	out       io.Writer
	formatter *phuslu.Formatter
}

func newPhusluLogger(config Configuration) (*phusluLogger, error) {
	logger := &phusluLogger{
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

func (l *phusluLogger) WithFields(keyValues Fields) Logger {
	return &phusluLogger{
		logger: &logphuslu.Logger{
			Caller: 2,
		},
		context:   logphuslu.Fields(keyValues),
		formatter: phuslu.NewFormatter(),
	}
}

func (l *phusluLogger) WithConfig(cfg *phuslu.Config) *phusluLogger {
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
func (l *phusluLogger) withTimeFormat(formatStr string) {
	l.logger.TimeFormat = formatStr
}

// sets the style of the logger.
func (l *phusluLogger) withStyle(style string) {
	if style == phuslu.StylePretty {
		l.useConsoleWriter()
	} else if style == phuslu.StyleJSON {
		l.useJSONWriter()
	}
}

// SetLevel sets the log level of the logger.
func (l *phusluLogger) withLogLevel(level string) {
	l.logger.Level = logphuslu.ParseLevel(level)
}

// useConsoleWriter sets the logger to use a console writer.
func (l *phusluLogger) useConsoleWriter() {
	l.setWriter(&logphuslu.ConsoleWriter{
		Writer:    l.out,
		Formatter: l.formatter.Format,
	})
}

// useConsoleWriter sets the logger to use a console writer.
func (l *phusluLogger) useFileWrite(filename string) {
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
func (l *phusluLogger) useJSONWriter() {
	l.setWriter(logphuslu.IOWriter{Writer: l.out})
}

// setWriter sets the writer of the logger.
func (l *phusluLogger) setWriter(writer logphuslu.Writer) {
	l.logger.Writer = writer
}

func (l *phusluLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *phusluLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *phusluLogger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *phusluLogger) Info(msg string) {
	l.logger.Info().Msg(msg)

}

func (l *phusluLogger) Infoln(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *phusluLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *phusluLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *phusluLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *phusluLogger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *phusluLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}

func (l *phusluLogger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

func (l *phusluLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

func (l *phusluLogger) Panic(msg string) {
	l.logger.Panic().Msg(msg)
}

func (l *phusluLogger) GetDelegate() interface{} {
	return l.logger
}

func (l *phusluLogger) SetLogLevel(logLevel string) error {
	l.logger.Level = logphuslu.ParseLevel(logLevel)
	return nil
}
