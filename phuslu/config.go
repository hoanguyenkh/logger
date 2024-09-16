package phuslu

const (
	Reset   Color = "\x1b[0m"
	Red     Color = "\x1b[31m"
	Green   Color = "\x1b[32m"
	Yellow  Color = "\x1b[33m"
	Magenta Color = "\x1b[35m"
	Gray    Color = "\x1b[90m"

	// log levels.
	traceColor   = Magenta
	debugColor   = Yellow
	infoColor    = Green
	warnColor    = Yellow
	errorColor   = Red
	fatalColor   = Red
	panicColor   = Red
	defaultColor = Red
	traceLabel   = "TRACE"
	debugLabel   = "DEBUG"
	infoLabel    = "INFO"
	warnLabel    = "WARN"
	errorLabel   = "ERROR"
	fatalLabel   = "FATAL"
	panicLabel   = "PANIC"
	defaultLabel = "???"

	StylePretty = "pretty"
	StyleJSON   = "json"
)

type Config struct {
	TimeFormat string `mapstructure:"time-format"`
	LogLevel   string `mapstructure:"log-level"`
	Style      string `mapstructure:"style"`
}

func DefaultConfig() Config {
	return Config{
		TimeFormat: "2000-01-02",
		LogLevel:   "info",
		Style:      StylePretty,
	}
}

// Color is a string that holds the hex color code for the color.
type Color string

// Raw returns the raw color code.
func (c Color) Raw() string {
	return string(c)
}
