package controllers

import (
	"riot/cli/src/api"
	"riot/cli/src/messaging"
)

func SwitchOption(option int) string {
	switch option {
	case 1:
		summonerName := messaging.AskForSummonerName()
		summonerData := api.SearchSummonerByName(summonerName)

		return summonerData
	case 2:
		return "Under construction..."
	default:
		return "! Default !"
	}
}
