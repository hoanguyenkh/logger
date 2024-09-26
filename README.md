# logger

Flexible logger interface that supports both zap, logrus and phuslu

# Example
## Using Zap
```go
package main

import (
	"github.com/KyberNetwork/logger"
)

func main() {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "debug",
		EnableFile:       false,
	}
	log, err := logger.NewLogger(config, logger.LoggerBackendZap)
	if err != nil {
		panic(err)
	}
	log.Info("This is an info message using Zap logger")
	log.Debug("This is a debug message using Zap logger")
	log.Error("This is an error message using Zap logger")
}

```


## Using Phuslu

```go
package main

import (
	"github.com/KyberNetwork/logger"
)

func main() {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "debug",
		EnableFile:       false,
	}
	log, err := logger.NewLogger(config, logger.LoggerBackendPhuslu)
	if err != nil {
		panic(err)
	}
	log.Info("This is an info message using Phuslu logger")
	log.Debug("This is a debug message using Phuslu logger")
	log.Error("This is an error message using Phuslu logger")
}

```


# Benchmark
    go test -bench=. -run=xxx -benchmem

## The benchmark results

    BenchmarkZapLogger-11           72950010                15.71 ns/op           16 B/op          1 allocs/op
    BenchmarkLogrusLogger-11        81084960                14.39 ns/op           16 B/op          1 allocs/op
    BenchmarkPhusluLogger-11        497384498                2.395 ns/op           0 B/op          0 allocs/op
