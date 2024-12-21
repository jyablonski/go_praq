package main

import "fmt"

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {

	// inline comment

	/* whole
		ass
	 comment */
	const jacobs_plan = "pro"
	const basicPlanName = "Basic Plan"
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(messageStart, age, messageEnd)
	fmt.Println(averageOpenRate, displayMessage)
	fmt.Println("plan:", basicPlanName)

	fmt.Println(billingCost(jacobs_plan))
	fmt.Println(concat("jacob", "hi"))
}
