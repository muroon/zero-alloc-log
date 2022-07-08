package poollog

import "sync/atomic"

type testStream struct {
	writeCount uint64
}

func (s *testStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *testStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}

func doneFunc(b []byte) {}
