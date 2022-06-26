package zero_alloc_log

import (
	"io"
	"os"
)

var LogWriter io.Writer // ログの出力(使用時に外部から指定可能)

type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
)

var IsPool bool

// event
type event struct {
	buf       []byte // ログmsgなど出力内容
	level    Level // ログLevel
	done      func(msg string) // ログ出力用完了channel (Panic, Fatal用)
}

func (e event) write() {
	if e.done != nil {
		defer e.done(string(e.buf))
	}
	LogWriter.Write(e.buf)
}

func Debug(msg string) {
	e := &event{
		level: DebugLevel,
		buf: []byte(msg),
	}
	e.write()
}

func Info(msg string) {
	e := &event{
		level: InfoLevel,
		buf: []byte(msg),
	}
	e.write()
}

func Warn(msg string) {
	e := &event{
		level: WarnLevel,
		buf: []byte(msg),
	}
	e.write()
}

func Error(err error) {
	e := &event{
		level: WarnLevel,
		buf: []byte(err.Error()),
	}
	e.write()
}

func Fatal(msg string) {
	e := &event{
		level: FatalLevel,
		buf: []byte(msg),
		done: func(msg string) { os.Exit(1) },
	}
	e.write()
}

func Panic(msg string) {
	e := &event{
		level: PanicLevel,
		buf: []byte(msg),
		done: func(msg string) { panic(msg) },
	}
	e.write()
}
