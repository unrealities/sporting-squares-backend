package nflstats

import (
	"net/http"
)

// GetScores fetches current NFL game scores
func GetScores() error {
	_, err := http.Get(espnNFLScoresURL())
	return err
}

// TODO: switch to ESPN:
func espnNFLScoresURL() string {
	return "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
}
