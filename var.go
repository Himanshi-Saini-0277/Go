package main
import(
	"fmt"
)

func main(){
	fmt.Println("I am Himanshi")

	var name string = "Himanshi"
	var a = 10.1 // It can be string or can be number or anything
	var b bool = true
	var c int = 1234
	var d float64 = 192.168
	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const pi = 3.14
	fmt.Println(pi)

	person := "Himanshi"
	fmt.Println(person)

	var Age = 19 // Can be called to external files
	var age = 19 // Can only be called in internal file
	fmt.Println(Age)
	fmt.Println(age)
}