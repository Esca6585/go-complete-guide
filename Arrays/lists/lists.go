package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := []float64{10.99, 9.99, 45.99, 20.00}

	fmt.Println(prices)
	fmt.Println(productNames)

	featuresPrices := prices[0:3]
	fmt.Println(featuresPrices)

	prices = append(prices, 11.99, 4.99, 6.9)
	fmt.Println(prices)

	discountPrice := []float64{1.99, 7.99, 8.999, 6.77}
	prices = append(prices, discountPrice...)

	fmt.Println(prices)
}
