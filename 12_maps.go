package main

import (
	"fmt"
)

func main(){
	Grade := make(map[string] int)

	Grade["Himanshi"] = 100
	Grade["Suman"] = 80
	Grade["Pawan"] = 60

	fmt.Printf("Grade of Himanshi is %d\n", Grade["Himanshi"])

	Grade["Himanshi"] = 95
	fmt.Printf("Grade of Himanshi is %d\n", Grade["Himanshi"])

	delete(Grade, "Himanshi")
	fmt.Printf("Grade of Himanshi is %d\n", Grade["Himanshi"])

	fmt.Printf("\n")

	grade, exist := Grade["Parshav"]
	fmt.Printf("Grade of Parshav is %d\n", grade)
	fmt.Println("Does Parshav exist ", exist)
	fmt.Printf("\n")

	grades, exists := Grade["Suman"]
	fmt.Printf("Grade of Suman is %d\n", grades)
	fmt.Println("Does Suman exist ", exists)
	fmt.Printf("\n")

	for index, value := range Grade {
		fmt.Printf("Key - %s, Value - %d\n", index, value)
	}
	
	fmt.Printf("\n")

	Person := map[string]int{
		"Himanshi": 100,
		"Suman": 80,
		"Pawan": 60,
	}
	for index, value := range Person {
		fmt.Printf("Key - %s, Value - %d\n", index, value)
	}

}