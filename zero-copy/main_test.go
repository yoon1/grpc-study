package main

import "testing"

const sourceFile = "source.txt"
const destinationWithoutZeroCopy = "destination_without_zero_copy.txt"
const destinationWithZeroCopy = "destination_with_zero_copy.txt"

// BenchmarkCopyWithoutZeroCopy
// BenchmarkCopyWithoutZeroCopy-10    	   21362	     59832 ns/op
func BenchmarkCopyWithoutZeroCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyWithoutZeroCopy(sourceFile, destinationWithoutZeroCopy)
	}
}
BenchmarkCopyWithZeroCopy
BenchmarkCopyWithZeroCopy-10    	   22272	     53925 ns/op

func BenchmarkCopyWithZeroCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyWithZeroCopy(sourceFile, destinationWithZeroCopy)
	}
}
