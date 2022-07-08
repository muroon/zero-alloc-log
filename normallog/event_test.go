package normallog

import "testing"

func TestNewEvent(t *testing.T) {
	stream := &testStream{}
	LogWriter = stream
	Info("test")
	t.Log(stream.writeCount)
}

func BenchmarkEvent(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("test")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkEventWithDone(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			InfoWithDone("test", doneFunc)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
