package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Summoner struct {
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	PUUID         string `json:"puuid"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

type SummonerStats struct {
	LeagueID     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

type ChallengerRanks struct {
	Entries []ChallengerEntry `json:"entries"`
}

type ChallengerEntry struct {
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Rank         string `json:"rank"`
}

type Mastery struct {
	ChampionId     int `json:"championId"`
	ChampionName   string
	ChampionLevel  int `json:"championLevel"`
	ChampionPoints int `json:"championPoints"`
}

type Champions struct {
	Champion []Champion `json:"data"`
}

type Champion struct {
	ID   string `json:"key"`
	Name string `json:"name"`
}

var riotApiKey, baseUrl string

func GetSummonerTopMastery(summonerName string) []Mastery {
	var summoner Summoner
	var topMastery []Mastery
	var champions Champions

	summoner = getSummonerGeneralInfo(
		summonerName,
		summoner,
	)

	requestUrl := fmt.Sprintf("%s/champion-mastery/v4/champion-masteries/by-summoner/%s/top", baseUrl, summoner.ID)

	client := &http.Client{}

	req, _ := http.NewRequest(
		"GET",
		requestUrl,
		nil,
	)

	req.Header.Set("X-Riot-Token", riotApiKey)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &topMastery)

	championsFile, _ := os.Open("champions.json")
	championsBytes, _ := io.ReadAll(championsFile)
	json.Unmarshal(championsBytes, &champions)

	defer championsFile.Close()

	for _, c := range champions.Champion {
		for i, t := range topMastery {
			if c.ID == strconv.FormatInt(int64(t.ChampionId), 10) {
				topMastery[i].ChampionName = c.Name
			}
		}
	}

	return topMastery
}

func GetChallengerQueue() ChallengerRanks {
	var challengerRanks ChallengerRanks

	requestUrl := fmt.Sprintf("%s/league/v4/challengerleagues/by-queue/RANKED_SOLO_5x5", baseUrl)

	client := &http.Client{}

	req, _ := http.NewRequest(
		"GET",
		requestUrl,
		nil,
	)

	req.Header.Set("X-Riot-Token", riotApiKey)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &challengerRanks)

	return challengerRanks
}

func SearchSummonerByName(summonerName string) []SummonerStats {
	var summoner Summoner
	var summonerStats []SummonerStats

	summoner = getSummonerGeneralInfo(
		summonerName,
		summoner,
	)

	summonerStats = getSummonerStats(
		summoner,
		summonerStats,
	)

	return summonerStats
}

func SetBaseValues(apiKey string) {
	baseUrl = "https://br1.api.riotgames.com/lol"
	riotApiKey = apiKey
}

func getSummonerGeneralInfo(summonerName string, summoner Summoner) Summoner {
	requestUrl := fmt.Sprintf(
		"%s/summoner/v4/summoners/by-name/%s",
		baseUrl,
		summonerName,
	)

	client := &http.Client{}

	req, _ := http.NewRequest(
		"GET",
		requestUrl,
		nil,
	)

	req.Header.Set("X-Riot-Token", riotApiKey)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &summoner)

	return summoner
}

func getSummonerStats(summoner Summoner, summonerStats []SummonerStats) []SummonerStats {
	requestUrl := fmt.Sprintf(
		"%s/league/v4/entries/by-summoner/%s",
		baseUrl,
		summoner.ID,
	)

	client := &http.Client{}

	req, _ := http.NewRequest(
		"GET",
		requestUrl,
		nil,
	)

	req.Header.Set("X-Riot-Token", riotApiKey)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal([]byte(body), &summonerStats)

	return summonerStats
}
