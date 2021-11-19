package nflstats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetScores fetches current NFL game scores
func GetScores() (Scores, error) {
	URL := espnNFLScoresURL()
	resp, err := http.Get(URL)
	if err != nil {
		return Scores{}, fmt.Errorf("nflStats#GetScores: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Scores{}, fmt.Errorf("nflStats#GetScores: reading Get response body error: %s", err)
	}

	espnNFLScoresResp := Scores{}
	err = json.Unmarshal(body, &espnNFLScoresResp)
	if err != nil {
		return Scores{}, fmt.Errorf("nflStats#GetScores: unmarshal error: %s", err)
	}

	return espnNFLScoresResp, nil
}

// TODO: switch to ESPN:
func espnNFLScoresURL() string {
	return "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
}
