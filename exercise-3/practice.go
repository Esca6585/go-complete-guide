package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) store() {
	file, _ := os.Create(prod.id)

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)

	file.WriteString(content)

	file.Close()
}

func (prod *Product) prinData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{
		id,
		t,
		d,
		p,
	}
}

func main() {

	createdProduct := getProduct()
	createdProduct.prinData()

	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter product data.")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
