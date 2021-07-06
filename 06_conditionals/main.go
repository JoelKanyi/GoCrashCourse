package main

import "fmt"

func main() {
	x:= 5
	y:= 10

	//If..else
	if x < y{
		fmt.Printf("%d is less than %d\n", x,y)
	}else {
		fmt.Printf("%d is less than %d\n", y,x)
	}

	//else if
	color:= "blue"

	if color == "blue" {
		fmt.Println("Color is blue")
	}else if color == "red"{
		fmt.Println("Color is red")
	}else{
		fmt.Println("Color not blue or red")
	}

	//Switch
	switch color {
	case "blue":
		fmt.Println("Color is Blue")
	case "green":
		fmt.Println("Color is green")
	default:
		fmt.Println("Color does not exist")
	}
}
