package main

//import "fmt"

func Solution(A []int, K int) []int {

	size := len(A)
	r := make([]int, size)

	if size < 1 {
		return A
	}

	for i := 0; i < size; i++ {
		r[(i+K)%size] = A[i]
	}

	return r
}
