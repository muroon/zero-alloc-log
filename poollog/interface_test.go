package poollog

import "testing"

func BenchmarkEventWithInterface(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			InfoByInterface("test")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
