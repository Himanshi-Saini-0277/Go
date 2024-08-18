package main

import(
	"fmt"
	"net/url"
)

func main(){
	myurl := "https://github.com/Himanshi-Saini-0277"
	fmt.Printf("Type of myurl is %T\n",myurl)
	
	tourl,err := url.Parse(myurl)
	
	if err!= nil{
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Type of myurl is %T\n",tourl)

	fmt.Println("Url is:", tourl)
	fmt.Println("Scheme: ", tourl.Scheme)
	fmt.Println("Host: ", tourl.Host)
	fmt.Println("Path: ", tourl.Path)
	fmt.Println("RawQuery: ", tourl.RawQuery)


	// Modify
	tourl.Path = "Himanshi"
	tourl.RawQuery = "Himanshi_Saini"

	newurl := tourl.String()
	fmt.Println("New Url is:", newurl)
}