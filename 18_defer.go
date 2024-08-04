package main

import "fmt"

func main(){
	fmt.Println("I will run")
	defer fmt.Println("I will go to stack and print just before main function is about to end")
	defer fmt.Println("I will go to top of stack and come out accordingly")
	fmt.Println("I will run before defer statements")
}