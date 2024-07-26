package main
import(
	"fmt"
)

func add (a, b int)(int) {
	return a+b
}

func multiply (a, b int)(result int) {
	result = a*b
	return
}

func main(){
	fmt.Println("I am Himanshi")
	ans := add(5,3)
	fmt.Println("Sum is:", ans)
	mul := multiply(5,3)
	fmt.Println("Product is:", mul)
}