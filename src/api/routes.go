package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var riotApiKey, baseUrl string

func SearchSummonerByName(summonerName string) string {
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

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func SetBaseValues() {
	baseUrl = "https://br1.api.riotgames.com/lol"
	riotApiKey = "RGAPI-5a0482c2-c35d-4dc9-9694-41bc17ff74aa"
}
