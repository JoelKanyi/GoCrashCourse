package main

import "fmt"

func main() {
	//Define map
	emails := make(map[string]string)

	//Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Ken"] = "ken@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	//Delete One
	delete(emails, "Ken")
	fmt.Println(emails)

	//Declare and add
	teams := map[string]string{"Chelsea": "chelsea fc", "Tottenham": "tottenham fc"}
	fmt.Println(teams)
}
