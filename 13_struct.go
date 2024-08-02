package main
import "fmt"

func main(){
	type Person struct{
		Name string
		ID int
		Age int
	}
	
	type Contact struct{
		Email string
		Phone int
	}
	
	type Address struct{
		House int
		Area string
		State string
	}
	
	type Employee struct{
		About Person
		Con Contact
		Addr Address
	}
	
	var Per1 Person
	Per1.Name = "Himanshi"
	Per1.ID = 11222553
	Per1.Age = 19
	
	var Emp1 Employee
	
	Emp1.About = Person{
		Name: "Parshav",
		ID: 11222543,
		Age: 18,
	}
	
	Emp1.Con.Email = "abc@gmail.com"
	Emp1.Con.Phone = 1234567890
	
	Emp1.Addr = Address{
		House: 123,
		Area: "Ambala",
		State: "Haryana",
	}
	
	fmt.Println(Per1)
	fmt.Println(Emp1)
}
