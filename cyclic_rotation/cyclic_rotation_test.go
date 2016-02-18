package main

import "testing"

func ValidateSolution(a []int, shift int, out []int, t *testing.T) (bool, []int) {

	result := Solution(a, shift)

	if len(result) != len(out) {
		return false, result
	}

	for i, _ := range result {
		if result[i] != out[i] {
			return false, result
		}
	}

	return true, result
}

func Assert(a []int, shift int, out []int, t *testing.T) {
	ok, result := ValidateSolution(a, shift, out, t)
	if !ok {
		t.Error("Fail to rotate", shift, " times. Expected: ", out, "Result:", result)
	}
}

func TestSolution(t *testing.T) {
	Assert([]int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}, t)
	Assert([]int{1, 2, 3, 4}, 2, []int{3, 4, 1, 2}, t)
}
