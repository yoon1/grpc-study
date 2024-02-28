package main

import "strconv"

func withZeroAllocation() {
	array := [1000000]int{}
	for i := 0; i < 1000000; i++ {
		array[i] = i
	}
}

func withoutZeroAllocation() {
	var slice []int
	for i := 0; i < 1000000; i++ {
		slice = append(slice, i)
	}
}

func generateString(i int) string {
	return "Number: " + strconv.Itoa(i) + "\n"
}

func generateStringZeroAlloc(i int) []byte {
	// zero allocation을 위해 직접 슬라이스에 값을 쓰기
	return append([]byte("Number: "), strconv.Itoa(i)...)
}

func main() {
	withoutZeroAllocation()
	withZeroAllocation()
}
