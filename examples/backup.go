package main

import (
	"fmt"
	"time"
)

func main() {
	start_timestamp := time.Now().UTC()

	// you just pass in the timestmp you want instead of yyyy-mm-dd ????????
	fmt.Println("Hello World at", start_timestamp.Format("2006-01-02 15:04:05 UTC"))

	end_timestamp := time.Now()

	execution_time := end_timestamp.Sub(start_timestamp)
	fmt.Println("Execution Time:", execution_time)

	numbers_divisble_by_3 := []int{}

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i%3 == 0 {
			numbers_divisble_by_3 = append(numbers_divisble_by_3, i)
			fmt.Println(i, "is divisible by 3")
		}

		if i == 9 {
			fmt.Println("Exiting mfer !")
			fmt.Println("Final List of Numbers Divisible by 3:", numbers_divisble_by_3)
		}
	}

	str1 := "test123"
	num_chars := return_str_chars(str1)

	if num_chars < 10 {
		fmt.Println("String is less than 10 characters")
	} else {
		fmt.Println("String is greater than or equal to 10 characters")
	}
}
