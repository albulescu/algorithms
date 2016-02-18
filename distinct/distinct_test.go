package main

import "testing"

func TestDistinct(t *testing.T) {
	distinct := Solution([]int{1, 2, 1, 3, 4, 2, 3})

	if distinct != 4 {
		t.Error("Distinct should be 4")
	}
}
