package main

import "fmt"

func main() {
	a := 5
	b := &a // assigning b to a pointer of a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // %T checks the type

	// use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
