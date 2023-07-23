package controllers

import (
	"fmt"
	"riot/cli/src/api"
	"riot/cli/src/messaging"
)

func SwitchOption(option int) {
	switch option {
	case 1:
		summonerName := messaging.AskForSummonerName()
		summonerStats := api.SearchSummonerByName(summonerName)

		for _, summoner := range summonerStats {
			fmt.Println("---------------------")
			fmt.Println("Queue Type:", summoner.QueueType)
			fmt.Println("Tier:", summoner.Tier)
			fmt.Println("Rank:", summoner.Rank)
			fmt.Println("League Points:", summoner.LeaguePoints)
			fmt.Println("Wins:", summoner.Wins)
			fmt.Println("Losses:", summoner.Losses)
		}

		return
	case 2:
		// Buscar lista de jogadores atuais no Challenger
		challengerEntries := api.GetChallengerQueue()

		for _, summoner := range challengerEntries.Entries {
			fmt.Println("---------------------")
			fmt.Println("Name:", summoner.SummonerName)
			fmt.Println("Rank:", summoner.Rank)
			fmt.Println("League Points:", summoner.LeaguePoints)
			fmt.Println("Wins:", summoner.Wins)
			fmt.Println("Losses:", summoner.Losses)
		}
	default:
		return
	}
}
