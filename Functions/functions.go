package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumbers()

	sum := add(a, b)

	print(sum)

}

func generateRandomNumbers() (int, int) {
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(10)
	return randomNumber1, randomNumber2
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func print(number int) {
	fmt.Printf("The Number is %v: ", number)
}
