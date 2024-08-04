package main

import (
	"fmt"
	"strconv"
)

func main(){
	a := 12
	fmt.Printf("%T\n", a)
	str := strconv.Itoa(a)
	fmt.Printf("%T\n", str)

	b := "12"
	fmt.Printf("%T\n", b)
	st, _ := strconv.Atoi(b)
	fmt.Printf("%T\n", st)

	c := "3.14"
	fmt.Printf("%T\n", c)
	strin, _ := strconv.ParseFloat(c, 64)
	fmt.Printf("%T\n", strin)

	d := "Himanshi"
	fmt.Printf("%T\n", d)
	stri, _ := strconv.ParseBool(d)
	fmt.Printf("%T\n", stri)
}
