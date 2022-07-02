package zero_alloc_log

type eventInterface interface {
	write()
}

func newEventInterface(buf []byte, level Level, done func(b []byte)) eventInterface {
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

func putEventInterface(e eventInterface) {
	if !isPool() {
		return
	}
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
