package main

import (
	"io"
	"os"
	"time"
)

func copyWithoutZeroCopy(sourceFile, destinationFile string) {
	source, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		panic(err)
	}
	defer destination.Close()

	// Using a buffer to demonstrate a non-zero-copy operation
	bufferSize := 1024
	buffer := make([]byte, bufferSize)

	for {
		// Read data into the buffer
		n, err := source.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// Write data from the buffer
		_, err = destination.Write(buffer[:n])
		if err != nil {
			panic(err)
		}
	}
}

func copyWithZeroCopy(sourceFile, destinationFile string) {
	source, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		panic(err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
}

func duration(msg string, f func()) {
	startTime := time.Now()
	f()
	elapsedTime := time.Since(startTime)
	println(msg, elapsedTime)
}

func main() {
	duration("Copy With Zero-Copy elapsed time:", func() { copyWithZeroCopy("source.txt", "destination_with_zero_copy.txt") })
	duration("Copy Without Zero-Copy elapsed time:", func() { copyWithoutZeroCopy("source.txt", "destination_without_zero_copy.txt") })
}
