package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()

	if plan == planPro {
		return allMessages[:], nil
	}

	if plan == planFree {
		return allMessages[0:2], nil
	}

	return nil, errors.New("unsupported Plan")
}

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}

	return costs
}

// can call this like total := sum(1.0, 2.1, 3.2)
// can take any number of arbitray ints and do what it needs to do
func sum(nums ...float64) float64 {
	total := 0.0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}

type cost struct {
	day   int
	value float64
}

func getCostsByDayt(costs []cost) []float64 {
	costsByDay := []float64{}

	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}

		return -1
	}
}

func main() {
	fmt.Println(getMessageWithRetries())

	mySlice := make([]int, 5)
	mySlice2 := []int{1, 1, 1, 1, 2}
	mySlice3 := make(map[string]string)

	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)

	matrix1 := createMatrix(5, 7)

	fmt.Println(matrix1)
}
