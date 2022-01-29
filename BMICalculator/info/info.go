package info

import "fmt"

const MainTitle = "BMI Calculator"
const Seperator = "------------------"
const WeightPromt = "Please enter your weight (kg): "
const HeightPromt = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(MainTitle)
	fmt.Println(Seperator)
}
