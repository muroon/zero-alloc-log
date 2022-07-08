package normallog

type eventInterface interface {
	write()
}

func newEventInterface(buf []byte, level Level, done func(b []byte)) eventInterface {
	return newEvent(buf, level, done)
}

// InfoByInterface is the func for bench test of using interface.
func InfoByInterface(msg string) {
	e := newEventInterface([]byte(msg), InfoLevel, nil)
	e.write()
}
