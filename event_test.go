package zero_alloc_log

import "testing"

func TestNewEventNormal(t *testing.T) {
	stream := &blackholeStream{}
	LogWriter = stream
	LogMode = ModeNormal
	Info("test")
	t.Log(stream.writeCount)
}

func BenchmarkEventNormal(b *testing.B) {
	stream := &blackholeStream{}
	LogWriter = stream
	LogMode = ModeNormal
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

func BenchmarkEventZeroAllocation(b *testing.B) {
	stream := &blackholeStream{}
	LogWriter = stream
	LogMode = ModeZeroAllocation
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

func BenchmarkEventWithDoneZeroAllocation(b *testing.B) {
	stream := &blackholeStream{}
	LogWriter = stream
	LogMode = ModeZeroAllocation
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DoneDebug("test", doneFunc)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
