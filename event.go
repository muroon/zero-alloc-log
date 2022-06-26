package zero_alloc_log

import (
	"io"
	"os"
	"sync"
)

var LogWriter io.Writer // ログの出力(使用時に外部から指定可能)
var LogMode Mode

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

type Mode int8

const (
	ModeNormal Mode = iota

	ModeZeroAllocation = iota
)

// event
type event struct {
	buf   []byte         // ログmsgなど出力内容
	level Level          // ログLevel
	done  func(b []byte) // ログ出力用完了channel (Panic, Fatal用)
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

func isPool() bool {
	return LogMode != ModeNormal
}

func newEvent(buf []byte, level Level, done func(b []byte)) *event {
	var e *event
	if isPool() {
		e = eventPool.Get().(*event)
	} else {
		e = &event{buf: make([]byte, 0, 500)}
	}
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
	if !isPool() {
		return
	}
	eventPool.Put(e)
}

func Debug(msg string) {
	e := newEvent([]byte(msg), DebugLevel, nil)
	e.write()
	putEvent(e)
}

func Info(msg string) {
	e := newEvent([]byte(msg), InfoLevel, nil)
	e.write()
	putEvent(e)
}

func Warn(msg string) {
	e := newEvent([]byte(msg), WarnLevel, nil)
	e.write()
	putEvent(e)
}

func Error(err error) {
	e := newEvent([]byte(err.Error()), ErrorLevel, nil)
	e.write()
	putEvent(e)
}

func Fatal(err error) {
	e := newEvent([]byte(err.Error()), FatalLevel, func(b []byte) { os.Exit(1) })
	e.write()
	putEvent(e)
}

func Panic(err error) {
	e := newEvent([]byte(err.Error()), PanicLevel, func(b []byte) { panic(b) })
	e.write()
	putEvent(e)
}

func DoneDebug(msg string, done func(b []byte)) {
	e := newEvent([]byte(msg), DebugLevel, done)
	e.write()
	putEvent(e)
}
