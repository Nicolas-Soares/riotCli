package messaging

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func GetOptionSelected() int {
	var optionSelect int

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for optionSelect == 0 || optionSelect > 3 {
		fmt.Printf("%s\n", yellow("=== LEAGUE OF LEGENDS CLI ==="))
		fmt.Println("What do you want to do?")
		fmt.Printf(`
%s - Search rank by summoner name
%s - See actual Challenger queue
%s - Exit CLI
		`,
			green("1"),
			green("2"),
			green("3"),
		)
		fmt.Println()

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
