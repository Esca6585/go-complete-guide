package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1:
	hobbies := [3]string{"relaxing", "working", "coding"}
	fmt.Println(hobbies)

	//2:
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3:
	mainhobbies := hobbies[:2]
	fmt.Println(mainhobbies)

	//4:
	fmt.Println(cap(mainhobbies))
	mainhobbies = mainhobbies[1:3]
	fmt.Println(mainhobbies)

	//5:
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	//6:
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics")
	fmt.Println(courseGoals)

	//7:
	products := []Product{
		{
			"first-product",
			"A First Product",
			129.99,
		},
		{
			"second-product",
			"A Second Product",
			135.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}
	products = append(products, newProduct)
	fmt.Println((products))
}
