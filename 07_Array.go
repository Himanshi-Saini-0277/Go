package main

import (
	"fmt"
)

func main() {
	var name[5] string
	name[0] = "Himanshi"
	name[2] = "Saini"

	fmt.Println(name)

	var numbers = [5] int {1,2,3,4,5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(numbers[2])

	var a[5] int
	fmt.Println(a)

	var b[5] bool
	fmt.Println(b)

	var c[5] string
	fmt.Println(c)
	fmt.Printf("%q ",c)
}
