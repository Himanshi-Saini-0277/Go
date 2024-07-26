package main
import(
	"fmt"
)

func divide (a, b float64)(float64, error) {
	if(b == 0){
		return a/b, fmt.Errorf("Cannot divide by 0")
	}
	return a/b, nil
}

func div (a, b float64)(float64, string) {
	if(b == 0){
		return a/b, "Cannot divide by 0"
	}
	return a/b, "nil"
}

func main(){
	ans, err := divide(10,0)
	fmt.Println("Quotiant is:", ans)
	if(err!= nil){
		fmt.Println("Error Handled")
	}
	answer, _ := div(10,4)
	fmt.Println("Quotiant is:", answer)
}