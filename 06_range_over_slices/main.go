package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// This program has two problems.
	// The first problem is the output reports the 1st prime
	// as the 0 th prime.
	// The second problem is the suffix the 1st prime should be
	// 1 st, not 1 th.
	// Can you fix both of these problems.

	for i, p := range primes {
		var suffix string

		switch i {
		case 0:
			suffix = "st"
		case 1:
			suffix = "nd"
		case 2:
			suffix = "rd"
		default:
			suffix = "th"
		}

		fmt.Println("The", i + 1, suffix, " prime is", p)
	}
}
