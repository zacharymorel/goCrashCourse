package main

import (
	"fmt"
)

// 			  var  type 	return type
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Zach"))
	fmt.Println(getSum(2, 3))
}
