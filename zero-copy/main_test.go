package main

import "testing"

const sourceFile = "source.txt"
const destinationWithoutZeroCopy = "destination_without_zero_copy.txt"
const destinationWithZeroCopy = "destination_with_zero_copy.txt"

func BenchmarkCopyWithoutZeroCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyWithoutZeroCopy(sourceFile, destinationWithoutZeroCopy)
	}
}

func BenchmarkCopyWithZeroCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyWithZeroCopy(sourceFile, destinationWithZeroCopy)
	}
}
