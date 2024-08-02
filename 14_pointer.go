package main
import "fmt"

func modify(val *int){
	*val = *val + 10
}

func main(){
	num := 2
	var ptr *int
	ptr = &num

	fmt.Println(ptr)
	fmt.Println(*ptr)

	name:= "Himanshi"
	ptr2 := &name

	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	var pointer *int
	if pointer == nil{
		fmt.Println("Empty Pointer")
	}

	value := 20
	modify(&value)
	fmt.Println(value)
}