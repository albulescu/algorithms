package main

func Solution(A []int) int {

	distinct := 0

	m := make(map[int]bool, 0)

	for _, n := range A {
		m[n] = true
	}

	for _ = range m {
		distinct++
	}

	return distinct
}
