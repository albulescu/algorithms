package main

import "testing"

func Assert(n int, out int, t *testing.T) {
	result := Solution(n)
	if result != out {
		t.Error(n, "should have a gap of", out, "not", result)
	}
}

func TestSolution(t *testing.T) {
	Assert(1041, 5, t)
	Assert(6, 0, t)
	Assert(328, 2, t)
	Assert(51712, 2, t)
	Assert(20, 1, t)
	Assert(74901729, 4, t)
	Assert(9, 2, t)
	Assert(90000000220000000, 4, t)
}
