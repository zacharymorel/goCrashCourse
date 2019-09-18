package main

import "fmt"

func main() {
	ids := []int{33, 67, 38, 98, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID : %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	} // _ (underscore) is used as a placeholder so we can get past the requirement to use all vars

	// Add add id's together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Wizz": "wizz@gmail.com", "Foo": "foo@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
