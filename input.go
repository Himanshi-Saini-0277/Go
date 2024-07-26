package main
import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("Hey, What's your name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hi, my name is:", name)

	fmt.Println("Hey, What's your name?")
	a := bufio.NewReader(os.Stdin)
	me, _ := a.ReadString('\n')
	fmt.Println("Hi, my name is:", me)

}