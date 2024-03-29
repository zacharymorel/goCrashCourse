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
	fmt.Println(len(fruitArrSlice)) // len() counts the number of items in the array/ slice
	fmt.Println(fruitArrSlice[1])
	fmt.Println(fruitArrSlice[1:2]) // we get a range
}
