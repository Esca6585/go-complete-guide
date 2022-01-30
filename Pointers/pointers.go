package main

import "fmt"

func main() {
	age := 24

	fmt.Println(age)

	myAge := &age

	fmt.Println(*myAge)
	fmt.Println(age)
}
