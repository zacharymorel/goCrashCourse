package main

import (
	"fmt"
	"math"

	"github.com/zacharymorel/goCrashCourse/03_packages/stringUtl"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(stringUtl.Reverse("olleh"))
}
