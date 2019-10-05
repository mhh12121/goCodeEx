package testRelated

import "testing"

var x uint64 = 10042424424242

func BenchmarkPopcount(b *testing.B) {
	// var x uint64
	// x = 10
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}

}
func BenchmarkPopcountLoop(b *testing.B) {
	// var x uint64
	// x = 10

	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}

}
func BenchmarkPopcountZeroBit(b *testing.B) {
	// var x uint64
	// x = 10

	for i := 0; i < b.N; i++ {
		PopCountLastNotZeroBit(x)
	}

}
func BenchmarkPopcountShift(b *testing.B) {
	// var x uint64
	// x = 10

	for i := 0; i < b.N; i++ {
		PopCountShift(x)
	}

}
