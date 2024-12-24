package main

import (
	"fmt"
	"strconv"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}

	return totalCost
}

func maxMessages(thresh float64) int {
	totalCost := 0.0

	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

func fizzbuzz() {
	// chars := make([]string, 100)

	for i := 1; i < 101; i++ {
		output := ""

		if i%3 == 0 {
			output += "fizz"
		}

		if i%5 == 0 {
			output += "buzz"
		}

		if output == "" {
			// this fucking converts the int to a unicode character instead of its string representation
			// output = string(i)
			output = strconv.Itoa(i)
		}

		fmt.Println(i, output)
	}
}

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {

		// 2 is the first special case, print that mf
		if n == 2 {
			fmt.Println(n)
			continue
		}

		// skip even numbers
		if n%2 == 0 {
			continue
		}
		isPrime := true

		// check if n is divisible by any number up to its square root
		// don't need to check past the square root because anything higher than sq rt has no chance
		// of multiplying evenly into n
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}

		fmt.Println(n)
	}
}

func main() {
	fizzbuzz()

	printPrimes(121)
}
