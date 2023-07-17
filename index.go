package main

import (
	"fmt"
	"riot/cli/src/api"
	"riot/cli/src/controllers"
	"riot/cli/src/messaging"
)

func main() {
	var optionSelected int
	var finalResult string

	optionSelected = messaging.GetOptionSelected()

	api.SetBaseValues()

	finalResult = controllers.SwitchOption(optionSelected)

	fmt.Println(finalResult)
}
