package main

import "fmt"

func main() {
	//Arrays
	var fruits [2]string

	//Assign values
	fruits[0] = "Mangoes"
	fruits[1] = "Banana"

	fmt.Println(fruits)
	fmt.Println(fruits[1])

	//Declare and assign
	counties := [2]string{"Narok", "Nakuru"}
	fmt.Println(counties)

	//Slices
	countries := []string{"Kenya", "Uganda", "Tanzania"}
	fmt.Println(countries)
	fmt.Println(len(countries))
	fmt.Println(countries[1:2])

}
