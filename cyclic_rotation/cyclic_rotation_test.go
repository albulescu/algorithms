package main

import "testing"

func AssertCyclicRotation(a []int, shift int, out []int, t *testing.T) bool {

	result := CyclicRotation(a, shift)

	if len(result) != len(out) {
		return false
	}

	for i := range result {
		if result[i] != out[i] {
			return false
		}
	}

	return true
}

func TestCyclicRotation(t *testing.T) {
	AssertCyclicRotation([]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 4, 5}, t)
}
