package main

// range lets you iterate over an array
// we're ignoring index value by using _ blank identifier

// numbers [5]int is an array w/ fixed size
// numbers []int is a slice with variable size

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Go can let you write variadic functions (...) that can take a variable number of arguments.
// However, you can use the append function which takes a slice and a new value, then returns
// a new slice with all the items in it.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {

		// check if numbers is empty and append 0 to sums
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
