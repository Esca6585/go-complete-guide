package main

import "fmt"

func main() {
	pi := 3.14
	// var pi float32 = 3.14
	radius := 5
	var circumference float64 = 2 * float64(pi) * float64(radius)

	fmt.Printf("For a radius of %v, the circle cirumference is %v (Type:%T)", radius, circumference, circumference)

	fmt.Println(circumference)

}
