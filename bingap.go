package main

import "strconv"

func Solution(N int) int {

	m := make(map[int]int, 0)
	bins := strconv.FormatInt(int64(N), 2)
	var pos int = 0
	var max int = 0
	for i, c := range bins {
		if string(c) == "0" {
			if _, ok := m[pos]; !ok {
				m[pos] = 0
			}
			m[pos]++
		} else {
			if m[pos] > max && i > 0 {
				max = m[pos]
			}
			pos = i + 1
		}
	}
	return max
}
