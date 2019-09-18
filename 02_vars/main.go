package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64 (length of intergers)
	// uint uint8 uint16 uint32 uint64 uintptr (can't be negative)
	// byte - alias for unit8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128 (really large numbers)

	// using var
	var name string = "Zach"
	var age int32 = 30
	const isCool = true

	// short hand
	name2 := "Zach2"
	size := 1.3

	// more shorthand
	nm, em := "Zach", "zach@gmail.com"

	fmt.Println(name, age)
	fmt.Println(name2)
	fmt.Printf("%T\n", size)
	fmt.Println(nm, em)
}
