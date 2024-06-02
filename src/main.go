package main

import "fmt"

func Color(color string) string {
	switch color {
	case "red", "yellow", "orange":
		return "warm"
	case "blue", "green", "purple":
		return "cool"
	default:
		return "unknown"
	}
}

func main() {
	chosen_color := "red"
	user_color_output := Color(chosen_color)

	return_text := fmt.Sprintf("The tone your color %s gives off is %s", chosen_color, user_color_output)

	fmt.Println(return_text)

}
