package main

import "fmt"

func main() {
	fmt.Println(greetings("Joel"))
	fmt.Println(addTwoNums(2, 4))
}

func greetings(name string) string {
	return "Hello " + name
}

func addTwoNums(num1, num2 int) int {
	return num1 + num2
}
