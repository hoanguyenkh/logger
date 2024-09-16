package logger_test

import (
	"github.com/KyberNetwork/logger"
	"github.com/stretchr/testify/require"
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
		{"PhusluLogger", logger.LoggerBackendPhuslu, false},
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

func TestPhuslu(t *testing.T) {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "debug",
		EnableFile:       false,
		FileJSONFormat:   false,
	}
	log, err := logger.NewLogger(config, logger.LoggerBackendPhuslu)
	require.NoError(t, err)

	log.Info("hoank")
	log.Debug("hoank")
	log.Error("hoank")
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
