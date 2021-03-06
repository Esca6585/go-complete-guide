package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/babamyrat1003/go-complete-guide/BMICalculator/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPromt)
	height = getUserInput(info.HeightPromt)

	return
}

func getUserInput(promptText string) (value float64) {

	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)

	return
}
