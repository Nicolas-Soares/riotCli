package main

import (
	"fmt"
	"riot/cli/src/api"
	"riot/cli/src/controllers"
	"riot/cli/src/messaging"
)

func main() {
	optionSelected := messaging.GetOptionSelected()

	api.SetBaseValues()
	finalResult := controllers.SwitchOption(optionSelected)

	fmt.Println(finalResult)
}
