package nflStats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: reduce GetESPNScores struct to app readable data
func GetScores() (Scores, error) {
	resp, err := GetESPNScores()
	for _, e := range resp.Events {
		fmt.Printf("Name: %+v \n", e.ShortName)
		for _, c := range e.Competitions {
			fmt.Printf("ID: %+v \n", c.ID)
		}
	}
	if err != nil {
		return Scores{}, fmt.Errorf("nflStats@GetScores: GetESPNScores, error: %s", err)
	}

	return Scores{}, nil
}

// GetESPNScores fetches current NFL game scores
func GetESPNScores() (EspnNFLScores, error) {
	URL := "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
	resp, err := http.Get(URL)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	espnNFLScoresResp := EspnNFLScores{}
	json.Unmarshal([]byte(body), &espnNFLScoresResp)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: reading Get response body error: %s", err)
	}

	return espnNFLScoresResp, nil
}
