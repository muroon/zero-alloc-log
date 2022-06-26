package zero_alloc_log

import "testing"

func TestNewEventNormal(t *testing.T) {
	stream := &blackholeStream{}
	LogWriter = stream
	Info("test")
	t.Log(stream.writeCount)
}

func BenchmarkEventNormal(b *testing.B) {
	stream := &blackholeStream{}
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
