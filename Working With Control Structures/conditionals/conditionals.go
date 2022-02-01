package conditionals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	userAge, err := getUserAge()
	if err != nil {
		panic(err)
	}
	if userAge >= 18 {
		fmt.Println("Welcome to the club")
	}
}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(strings.TrimSpace(userAgeInput), 0, 64)
	return int(userAge), err
}
