package main

import (
	"errors"
	"fmt"
	"os"
	"riot/cli/src/api"
	"riot/cli/src/controllers"
	"riot/cli/src/messaging"

	"github.com/joho/godotenv"
)

func main() {
	continueTask := 1

	setEnv()

	for continueTask == 1 {
		messaging.ClearTerminal()
		optionSelected := messaging.GetOptionSelected()

		if optionSelected == 3 {
			break
		}

		riotApiKey := os.Getenv("RIOT_API_KEY")
		api.SetBaseValues(riotApiKey)

		controllers.SwitchOption(optionSelected)

		fmt.Println("---------------------")
		fmt.Println()
		fmt.Println("Continue using the CLI?")
		fmt.Println(`
			1 - Yes
			2 - No
		`)

		fmt.Scan(&continueTask)
	}

	fmt.Println("=== Exiting Riot CLI ===")
}

func setEnv() {
	err := godotenv.Load()

	if err != nil {
		errors.New("Error setting env variables")
	}
}
