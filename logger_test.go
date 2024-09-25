package logger_test

import (
	"github.com/KyberNetwork/logger"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "info",
		EnableFile:       false,
		FileJSONFormat:   false,
	}

	tests := []struct {
		name    string
		backend logger.LoggerBackend
		wantErr bool
	}{
		{"ZapLogger", logger.LoggerBackendZap, false},
		{"LogrusLogger", logger.LoggerBackendLogrus, false},
		{"phusluLogger", logger.LoggerBackendPhuslu, false},
		{"InvalidLogger", logger.LoggerBackend(999), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := logger.NewLogger(config, tt.backend)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogger(t *testing.T) {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "debug",
		EnableFile:       false,
		FileJSONFormat:   false,
	}
	log, err := logger.NewLogger(config, logger.LoggerBackendPhuslu)
	require.NoError(t, err)

	log.Info("test log phuslu info")
	log.Debug("test log phuslu debug")
	log.Error("test log phuslu error")

	log, err = logger.NewLogger(config, logger.LoggerBackendZap)
	require.NoError(t, err)

	log.Info("test log zap info")
	log.Debug("test log zap debug")
	log.Error("test log zap error")

	log1 := logger.DefaultLogger()
	log1.Info("test default logger")
	log1.Error("test default logger")
}

func TestPhuslu_useFile(t *testing.T) {
	config := logger.Configuration{
		EnableConsole:    false,
		EnableJSONFormat: false,
		ConsoleLevel:     "debug",
		EnableFile:       true,
		FileJSONFormat:   false,
		FileLocation:     "test-logs/",
	}
	log, err := logger.NewLogger(config, logger.LoggerBackendPhuslu)
	require.NoError(t, err)

	log.Info("test info")
	log.Debug("test debug")
	log.Error("test error")
	// Check if file exists
	_, err = os.Stat(config.FileLocation)
	require.NoError(t, err, "log file should exist")

	// Remove the directory after test
	err = os.RemoveAll("test-logs")
	require.NoError(t, err, "failed to remove log directory")
}

func benchmarkLogger(b *testing.B, backend logger.LoggerBackend) {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "info",
		EnableFile:       false,
		FileJSONFormat:   false,
	}
	log, err := logger.NewLogger(config, backend)
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("benchmarking log message")
	}
}

func BenchmarkZapLogger(b *testing.B) {
	benchmarkLogger(b, logger.LoggerBackendZap)
}

func BenchmarkLogrusLogger(b *testing.B) {
	benchmarkLogger(b, logger.LoggerBackendLogrus)
}

func BenchmarkPhusluLogger(b *testing.B) {
	benchmarkLogger(b, logger.LoggerBackendPhuslu)
}
