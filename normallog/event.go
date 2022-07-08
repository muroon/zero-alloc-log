package normallog

import (
	"io"
	"os"
)

var LogWriter io.Writer

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

// event
type event struct {
	buf   []byte         // log message
	level Level          // log level
	done  func(b []byte) // after writing (for Panic, Fatal)
}

func (e event) write() {
	if e.done != nil {
		defer e.done(e.buf)
	}
	LogWriter.Write(e.buf)
}

func newEvent(buf []byte, level Level, done func(b []byte)) *event {
	return &event{
		buf:   buf,
		level: level,
		done:  done,
	}
}

// Debug writes Debug level log.
func Debug(msg string) {
	e := newEvent([]byte(msg), DebugLevel, nil)
	e.write()
}

// Info writes Info level log.
func Info(msg string) {
	e := newEvent([]byte(msg), InfoLevel, nil)
	e.write()
}

// Warn writes Warn level log.
func Warn(msg string) {
	e := newEvent([]byte(msg), WarnLevel, nil)
	e.write()
}

// Error writes Error level log.
func Error(err error) {
	e := newEvent([]byte(err.Error()), ErrorLevel, nil)
	e.write()
}

// Fatal writes Fatal level log.
func Fatal(err error) {
	e := newEvent([]byte(err.Error()), FatalLevel, fatalFunc)
	e.write()
}

func fatalFunc(b []byte) { os.Exit(1) }

// Panic writes Panic level log.
func Panic(err error) {
	e := newEvent([]byte(err.Error()), PanicLevel, panicFunc)
	e.write()
}

func panicFunc(b []byte) { panic(b) }

// InfoWithDone is the func for bench test of using channel.
func InfoWithDone(msg string, done func(b []byte)) {
	e := newEvent([]byte(msg), InfoLevel, done)
	e.write()
}
