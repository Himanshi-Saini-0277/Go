package main

import(
	"fmt"
	"strings"
)

func main(){
	data := "    one two    three four    "

	a := strings.Split(data," ")
	b := strings.Count(data, "two")
	c := strings.TrimSpace(data)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
