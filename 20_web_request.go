package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!= nil {
		fmt.Println("Error in getting.")
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response is %T\n", res)

	read, er := ioutil.ReadAll(res.Body)
	if er!= nil {
		fmt.Println("Error in reading ", er)
		return
	}
	fmt.Println("Response is: ",string(read))
}