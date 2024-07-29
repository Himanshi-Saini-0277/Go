package main

import(
	"fmt"
)

func main(){
	age := 19

	if age<3 {
		fmt.Println("You are a Baby")
	} else if (age<40 && age>2){
		fmt.Println("You are a Young Adult")
	} else if (age<60 && age>39){
		fmt.Println("You are a Middle-aged Adult")
	} else{
		fmt.Println("You are an Old Adult")
	}
}