package main

func CyclicRotation(A []int, K int) []int {
	r := make([]int, 0)
	for i, e := range A {
		r[i] = e
	}
	return r
}
