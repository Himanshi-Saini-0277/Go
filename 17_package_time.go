package main

import (
	"fmt"
	"time"
)

func main(){
	a := time.Now()
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := a.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println(b)

	c := a.Format("2006/01/02, Monday, 03:04:05 PM")
	fmt.Println(c)

	layout := "02-01-2006"
	data := "04-08-2024"
	formatted, _ := time.Parse(layout, data)
	fmt.Println(formatted)

	d := a.Add(24 * time.Hour)
	e := d.Format("02-01-2006, Monday")
	fmt.Println(e)
}