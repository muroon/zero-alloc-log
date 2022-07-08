package poollog

import (
	"io"
	"os"
	"sync"
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

var eventPool = &sync.Pool{
	New: func() interface{} {
		return &event{
			buf: make([]byte, 0, 500),
		}
	},
}

func newEvent(buf []byte, level Level, done func(b []byte)) *event {
	e := eventPool.Get().(*event)
	e.buf = e.buf[:0]
	e.buf = append(e.buf, buf...)
	e.level = level
	e.done = nil
	if done != nil {
		e.done = done
	}
	return e
}

func putEvent(e *event) {
	eventPool.Put(e)
}

// Debug writes Debug level log.
func Debug(msg string) {
	e := newEvent([]byte(msg), DebugLevel, nil)
	e.write()
	putEvent(e)
}

// Info writes Info level log.
func Info(msg string) {
	e := newEvent([]byte(msg), InfoLevel, nil)
	e.write()
	putEvent(e)
}

// Warn writes Warn level log.
func Warn(msg string) {
	e := newEvent([]byte(msg), WarnLevel, nil)
	e.write()
	putEvent(e)
}

// Error writes Error level log.
func Error(err error) {
	e := newEvent([]byte(err.Error()), ErrorLevel, nil)
	e.write()
	putEvent(e)
}

// Fatal writes Fatal level log.
func Fatal(err error) {
	e := newEvent([]byte(err.Error()), FatalLevel, fatalFunc)
	e.write()
	putEvent(e)
}

func fatalFunc(b []byte) { os.Exit(1) }

// Panic writes Panic level log.
func Panic(err error) {
	e := newEvent([]byte(err.Error()), PanicLevel, panicFunc)
	e.write()
	putEvent(e)
}

func panicFunc(b []byte) { panic(b) }

// InfoWithDone is the func for bench test of using channel.
func InfoWithDone(msg string, done func(b []byte)) {
	e := newEvent([]byte(msg), DebugLevel, done)
	e.write()
	putEvent(e)
}
