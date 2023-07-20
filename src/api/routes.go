package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

var riotApiKey, baseUrl string

func GetChallengerQueue() {}

func SearchSummonerByName(summonerName string) []SummonerStats {
	var summoner Summoner
	var summonerStats []SummonerStats
	var requestUrl string

	summoner = getSummonerGeneralInfo(
		requestUrl,
		summonerName,
		summoner,
	)

	summonerStats = getSummonerStats(
		requestUrl,
		summoner,
		summonerStats,
	)

	return summonerStats
}

func SetBaseValues(apiKey string) {
	baseUrl = "https://br1.api.riotgames.com/lol"
	riotApiKey = apiKey
}

func getSummonerGeneralInfo(requestUrl string, summonerName string, summoner Summoner) Summoner {
	requestUrl = fmt.Sprintf(
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

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &summoner)

	return summoner
}

func getSummonerStats(requestUrl string, summoner Summoner, summonerStats []SummonerStats) []SummonerStats {
	requestUrl = fmt.Sprintf(
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

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(body), &summonerStats)

	return summonerStats
}
