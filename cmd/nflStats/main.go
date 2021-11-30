package nflStats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: reduce GetESPNScores struct to app readable data
func GetScores() (Scores, error) {
	resp, err := GetESPNScores()
	if err != nil {
		return Scores{}, fmt.Errorf("nflStats@GetScores: GetESPNScores, error: %s", err)
	}
	fmt.Printf("main: GetESPNScores, resp: %+v", resp)

	return Scores{}, nil
}

// GetESPNScores fetches current NFL game scores
// TODO:
// nflStats@GetScores: GetESPNScores, error: nflStats#GetESPNScores: unmarshal error: json: cannot unmarshal number 7.0 into Go struct field .events.competitions.competitors.linescores.value of type int
func GetESPNScores() (EspnNFLScores, error) {
	URL := "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
	resp, err := http.Get(URL)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: reading Get response body error: %s", err)
	}

	espnNFLScoresResp := EspnNFLScores{}
	err = json.Unmarshal(body, &espnNFLScoresResp)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: unmarshal error: %s", err)
	}

	return EspnNFLScores{}, nil
}
