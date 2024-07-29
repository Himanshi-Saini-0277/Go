package main

import (
	"fmt"
)

func main() {
	var numbers = [] int {1,2,3,4,5}
	numbers = append (numbers, 6,7,8,9)
	fmt.Println(numbers)
	fmt.Println("Length is ", len(numbers))
	fmt.Println("Capacity is ", cap(numbers))
	fmt.Println(numbers[2])
	fmt.Printf("Type of data is %T\n", numbers)

	a := make([]int, 3, 6)
	a = append(a, 4)
	a = append(a, 5)
	a = append(a, 6)
	a = append(a, 7)
	fmt.Println(a)
	fmt.Println("Length is ", len(a))
	fmt.Println("Capacity is ", cap(a))
}