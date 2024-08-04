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

	str1 := "Himanshi"
	str2 := "Saini"

	d := strings.Join([]string{str1,str2}, " ")
	fmt.Println(d)
}
