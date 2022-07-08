package poollog

type eventInterface interface {
	write()
}

func newEventInterface(buf []byte, level Level, done func(b []byte)) eventInterface {
	return newEvent(buf, level, done)
}

func putEventInterface(e eventInterface) {
	if _, ok := e.(*event); !ok {
		panic("invalid type")
	}
	eventPool.Put(e)
}

// InfoByInterface is the func for bench test of using interface.
func InfoByInterface(msg string) {
	e := newEventInterface([]byte(msg), InfoLevel, nil)
	e.write()
	putEventInterface(e)
}
