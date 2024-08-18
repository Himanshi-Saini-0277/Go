package main

import(
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main(){
	person := Person{Name:"Himanshi", Age:19, IsAdult:true}
	fmt.Println("Person data is:", person)

	jsonData, err := json.Marshal(person)
	if err!= nil {
		fmt.Println("Error",err)
		return
	}
	fmt.Println("Person data after Marshal:", string(jsonData))

	var personData Person
	er := json.Unmarshal(jsonData, &personData)
	if er!= nil {
		fmt.Println("Error",er)
		return
	}
	fmt.Println("Person data after Unmarshal:", personData)
}