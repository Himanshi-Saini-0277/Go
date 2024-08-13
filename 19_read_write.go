package main
import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main(){
	file, err := os.Create("file.txt")
	if err!= nil{
		fmt.Println("Error while creating the file")
		return
	}

	content := "Hi I am Himanshi"
	b, errors := io.WriteString(file, content + "\n")
	fmt.Printf("byte written:", b)
	if errors != nil{
		fmt.Println("Error while writting in the file")
		return
	}

	file, erro := os.Open("file.txt")
	if erro != nil {
		fmt.Println("Error while opening the file")
	}

	buffer := make([]byte, 1024)
	for{
		n,er := file.Read(buffer)

		if er == io.EOF{
			break
		}
		if er != nil {
			fmt.Println("Error while reading the file")
			return
		}
		fmt.Println()
		fmt.Printf(string(buffer[:n]))
	}

	con, e := ioutil.ReadFile("file.txt")
	if e != nil {
		fmt.Println("Error while reading the file")
		return
	}
	fmt.Printf(string(con))
}