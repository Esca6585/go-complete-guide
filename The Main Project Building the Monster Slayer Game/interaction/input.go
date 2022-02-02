package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {

	for {
		getPlayerChoice, _ := getPlayerInput()

		if getPlayerChoice == "1" {
			return "ATTACK"
		} else if getPlayerChoice == "2" {
			return "HEAL"
		} else if getPlayerChoice == "3" && specialAttackIsAvailable {
			return "SPEACIAL_ATTACK"
		}

		fmt.Println("Fetching the user input failed. Please try again.")
	}

}

// This function Name is started with lowercase letter because it is used inside of this file
func getPlayerInput() (string, error) {

	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	userInput = strings.Replace(strings.TrimSpace(userInput), "\n", "", -1)

	return userInput, nil
}
