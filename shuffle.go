package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * Fisher–Yates shuffle
 * https://en.wikipedia.org/wiki/Fisher–Yates_shuffle
 */
func Shuffle(n []int) []int {

	i := len(n)

	for i > 1 {
		i = i - 1
		rand.Seed(time.Now().UTC().UnixNano())
		j := rand.Intn(i)
		n[j], n[i] = n[i], n[j]
	}

	return n
}

func main() {
	fmt.Println(Shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
