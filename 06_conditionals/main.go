package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y) // %d is base 10  values
	} else {
		fmt.Printf("%d is less than %d\n", y, x) // %d is base 10  values
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	// SWITCH
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}
