package main

import "fmt"

func main() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
			31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
			73, 79, 83, 89, 97, 101, 103, 107, 109,
			113, 127, 131, 137, 139, 149, 151, 157,
			163, 167, 173, 179, 181, 191, 193, 197, 199}
	fmt.Println(primes)

	// Write a program to print only the primes less than 10
	// loop through the slice of primes and test if the value
	// is less than 10. When you find a value that is 10 or more
	// slice the list of primes at that point and print it.
	for index, value := range primes {
		if value < 10 {
			fmt.Println(value)
		} else {
			fmt.Println(primes[:index])
		}
	}

	// Bonus: write a print only the two digit primes.
	for _, value := range primes {
		if value > 9 && value < 100 {
			fmt.Println(value)
		}
	}

	// see src/slices/slices9/main.go for the answer
}
