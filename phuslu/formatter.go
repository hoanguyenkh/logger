package phuslu

import (
	"io"

	"github.com/phuslu/log"
)

// Formatter is a custom formatter for log messages.
type Formatter struct {
	// KeyColors applies specific colors to log entries based on their keys.
	keyColors map[string]Color
	// KeyValColors applies specific colors to log entries based on their keys
	// and values.
	keyValColors map[string]Color
}

// NewFormatter creates a new Formatter with default settings.
func NewFormatter() *Formatter {
	return &Formatter{
		keyColors:    make(map[string]Color),
		keyValColors: make(map[string]Color),
	}
}

// Format formats the log message.
func (f *Formatter) Format(
	out io.Writer,
	args *log.FormatterArgs,
) (int, error) {
	buffer, ok := byteBufferPool.Get().(*byteBuffer)
	if !ok {
		panic("failed to get byte buffer from pool")
	}

	buffer.Reset()
	defer byteBufferPool.Put(buffer)

	var color, label string
	switch args.Level {
	case "trace":
		color, label = traceColor.Raw(), traceLabel
	case "debug":
		color, label = debugColor.Raw(), debugLabel
	case "info":
		color, label = infoColor.Raw(), infoLabel
	case "warn":
		color, label = warnColor.Raw(), warnLabel
	case "error":
		color, label = errorColor.Raw(), errorLabel
	case "fatal":
		color, label = fatalColor.Raw(), fatalLabel
	case "panic":
		color, label = panicColor.Raw(), panicLabel
	default:
		color, label = defaultColor.Raw(), defaultLabel
	}

	f.printWithColor(args, buffer, color, label)
	f.ensureLineBreak(buffer)

	if args.Stack != "" {
		buffer.Bytes = append(buffer.Bytes, args.Stack...)
		if args.Stack[len(args.Stack)-1] != '\n' {
			buffer.Bytes = append(buffer.Bytes, '\n')
		}
	}

	return out.Write(buffer.Bytes)
}

// AddKeyColor adds a key and color to the keyColors map.
func (f *Formatter) AddKeyColor(key string, color Color) {
	f.keyColors[key] = color
}

// AddKeyValColor adds a key and color to the keyValColors map.
func (f *Formatter) AddKeyValColor(key string, val string, color Color) {
	f.keyValColors[key+val] = color
}

// printWithColor prints the log message with color.
func (f *Formatter) printWithColor(
	args *log.FormatterArgs,
	b *byteBuffer,
	color, label string,
) {
	var (
		ok      bool
		kvColor Color
		kColor  Color
	)
	f.formatHeader(args, b, color, label)

	b.Bytes = append(b.Bytes, ' ')
	b.Bytes = append(b.Bytes, args.Message...)
	for _, kv := range args.KeyValues {
		b.Bytes = append(b.Bytes, ' ')
		// apply the key+value color if configured, otherwise apply key color
		if kvColor, ok = f.keyValColors[kv.Key+kv.Value]; ok {
			b.Bytes = append(b.Bytes, kvColor.Raw()...)
		} else if kColor, ok = f.keyColors[kv.Key]; ok {
			b.Bytes = append(b.Bytes, kColor.Raw()...)
		}
		b.Bytes = append(b.Bytes, kv.Key...)
		b.Bytes = append(b.Bytes, '=')
		b.Bytes = append(b.Bytes, kv.Value...)
		b.Bytes = append(b.Bytes, Reset...)
	}
}

// formatHeader formats the header of the log message.
func (f *Formatter) formatHeader(
	args *log.FormatterArgs,
	b *byteBuffer,
	color, label string,
) {
	b.Bytes = append(b.Bytes, Gray...)
	b.Bytes = append(b.Bytes, args.Time...)
	b.Bytes = append(b.Bytes, ' ')
	b.Bytes = append(b.Bytes, color...)
	b.Bytes = append(b.Bytes, label...)
	if args.Caller != "" {
		b.Bytes = append(b.Bytes, args.Goid...)
		b.Bytes = append(b.Bytes, ' ')
		b.Bytes = append(b.Bytes, args.Caller...)
		b.Bytes = append(b.Bytes, ' ')
	}
	b.Bytes = append(b.Bytes, Reset...)
}

// ensureLineBreak ensures the log message ends with a line break.
func (f *Formatter) ensureLineBreak(b *byteBuffer) {
	if b.Bytes == nil {
		b.Bytes = make([]byte, 0)
	}
	length := len(b.Bytes)
	if length == 0 || b.Bytes[length-1] != '\n' {
		b.Bytes = append(b.Bytes, '\n')
	}
}