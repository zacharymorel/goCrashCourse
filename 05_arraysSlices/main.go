package main

import "fmt"

func main() {
	//  - Arrays - //
	// var fruitArr [2]string

	// assign values
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"
	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Declare and Assign
	fruitArr := [2]string{"apple", "orange"}
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// - Slices - //
	fruitArrSlice := []string{"apple", "orange", "grape"}
	fmt.Println(fruitArrSlice)
	fmt.Println(len(fruitArrSlice))
	fmt.Println(fruitArrSlice[1])
}
