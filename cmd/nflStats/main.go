package nflStats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: reduce GetESPNScores struct to app readable data
func GetScores() (Scores, error) {
	_, err := GetESPNScores()
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

	// Looking at fields in response
	fmt.Printf("GetESPNScores resp.Body: %+v", resp.Body)
	// TODO: No Leagues
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	espnNFLScoresResp := EspnNFLScores{}
	json.NewDecoder(resp.Body).Decode(espnNFLScoresResp)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: reading Get response body error: %s", err)
	}

	return EspnNFLScores{}, nil
}
