package main

import(
	"fmt"
)

func main(){
	for i:=0 ; i<4; i++{
		fmt.Println(i)
	}
	fmt.Printf("\n")

	count := 0
	for {
		fmt.Println("Infinite loop")
		count ++

		if count == 3{
			break
		}
	}
	fmt.Printf("\n")

	number := [] int {1,2,3,4}
	for index, value := range number{
		fmt.Printf("%d at %d\n", value, index)
	}
	fmt.Printf("\n")

	name := "Himanshi"
	for index, char := range name{
		fmt.Printf("%c at %d\n", char, index)
	}
}