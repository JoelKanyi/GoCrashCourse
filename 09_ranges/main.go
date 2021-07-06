package main

import "fmt"

func main() {
	ids:= []int{33,56,87,90,34}

	//Loop through ids
	for i, id:= range ids{
		fmt.Printf("%d -ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids{
		fmt.Println(id)
	}

	//Add ids
	sum:= 0
	for _, i:= range ids{
		sum += i
	}
	fmt.Println("Sum= ", sum)

	//Range with map
	teams:= map[string]string{"Chelsea":"chelsea fc", "Tottenham":"tottenham fc"}

	for k,v := range teams{
		fmt.Printf("%s : %s\n", k,v )
	}
}
