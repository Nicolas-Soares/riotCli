package messaging

import (
	"fmt"
	"os"
	"os/exec"
)

func GetOptionSelected() int {
	var optionSelect int

	for optionSelect == 0 || optionSelect > 3 {
		fmt.Println("=== LEAGUE OF LEGENDS CLI ===")
		fmt.Println("What do you want to do?")
		fmt.Println(`
			1 - Search rank by summoner name
			2 - See actual Challenger queue
			3 - Exit CLI
		`)

		fmt.Scan(&optionSelect)

		ClearTerminal()
	}

	return optionSelect
}

func AskForSummonerName() string {
	var summonerName string

	fmt.Println("Type your summoner name: ")
	fmt.Scan(&summonerName)

	return summonerName
}

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
