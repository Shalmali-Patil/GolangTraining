package main

import "fmt"

func main() {
	fmt.Println(findLargestPrimeFactor(13195))        // Answer: 29
	fmt.Println(findLargestPrimeFactor(600851475143)) // Answer: 6857
}

func findLargestPrimeFactor(n int) int {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			n = n / i
			i = 2
		}
	}
	return n
}

/*
Project Euler question: Largest prime factor
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
