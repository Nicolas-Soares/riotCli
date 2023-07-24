package main

import (
	"errors"
	"fmt"
	"os"
	"riot/cli/src/api"
	"riot/cli/src/controllers"
	"riot/cli/src/messaging"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	// Set color patterns
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	continueTask := 1

	setEnv()

	for continueTask == 1 {
		messaging.ClearTerminal()
		optionSelected := messaging.GetOptionSelected()

		if optionSelected == 4 {
			break
		}

		riotApiKey := os.Getenv("RIOT_API_KEY")
		api.SetBaseValues(riotApiKey)

		controllers.SwitchOption(optionSelected)

		fmt.Println("---------------------")
		fmt.Println()
		fmt.Printf("%s", yellow("Continue using the CLI?"))
		fmt.Printf(`
%s - Yes
%s - No
		`,
			green("1"),
			green("2"),
		)
		fmt.Println()

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
