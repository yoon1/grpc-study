package main

import "testing"

func BenchmarkWithoutZeroAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutZeroAllocation()
	}
}

func BenchmarkWithZeroAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withZeroAllocation()
	}
}
