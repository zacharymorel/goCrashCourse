package main

import (
	"fmt"
)

func main() {
	// define map       key    val
	emails := make(map[string]string)

	// assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Wizz"] = "wizz@gmail.com"
	emails["Foo"] = "foo@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete
	delete(emails, "Foo")
	fmt.Println(emails)

	// assign when declaring
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Wizz": "wizz@gmail.com", "Foo": "foo@gmail.com"}

	fmt.Println(emails2)
	fmt.Println(len(emails2))
	fmt.Println(emails2["Bob"])

	// delete
	delete(emails2, "Foo")
	fmt.Println(emails2)
}
