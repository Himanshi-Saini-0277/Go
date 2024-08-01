package main

import (
	"fmt"
)

func main(){
	day := 3

	switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Unknown Day")
	}

	month := "May"

	switch month{
		case "November","December","January":
			fmt.Println("Winter")
		case "May", "June", "July":
			fmt.Println("Summer")
		default:
			fmt.Println("Other Season")
	}

	temperature := -10

	switch{
		case temperature < 0:
			fmt.Println("Freezing")
		case temperature >=0 && temperature <20:
			fmt.Println("Cold")
		case temperature >=20 && temperature <40:
			fmt.Println("Hot")
		default:
			fmt.Println("Melting")
	}
}
