package main

import "fmt"

func main() {
	a:= 15
	b:= &a

	fmt.Println(a,b)

	//Type
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)

	//Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change vsl with pointer
	*b = 10
	fmt.Println(a)
}
