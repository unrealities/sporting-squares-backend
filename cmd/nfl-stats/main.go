package nflstats

import (
	"fmt"
	"net/http"
)

// GetScores fetches current NFL game scores
func GetScores() error {
	URL := espnNFLScoresURL()
	resp, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("mlbStats#GetSchedule: Get %s, error: %s", URL, err)
	}
	defer resp.Body.Close()
	return err
}

// TODO: switch to ESPN:
func espnNFLScoresURL() string {
	return "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
}
