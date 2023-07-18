package main

import (
	"errors"
	"os"
	"riot/cli/src/api"
	"riot/cli/src/controllers"
	"riot/cli/src/messaging"

	"github.com/joho/godotenv"
)

func main() {
	var optionSelected int
	var riotApiKey string

	setEnv()

	optionSelected = messaging.GetOptionSelected()

	riotApiKey = os.Getenv("RIOT_API_KEY")
	api.SetBaseValues(riotApiKey)

	controllers.SwitchOption(optionSelected)
}

func setEnv() {
	err := godotenv.Load()

	if err != nil {
		errors.New("Error setting env variables")
	}
}
