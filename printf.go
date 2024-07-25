package main
import(
	"fmt"
)

func main(){
	fmt.Println("I am Himanshi")

	var name string = "Himanshi"
	var c int = 1234
	var d float64 = 9.16824
	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", c)
	fmt.Printf("%2f\n", d)

	// %T tells the type of data
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", name)
}