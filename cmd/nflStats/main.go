package nflStats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// TODO: reduce GetESPNScores struct to app readable data
func GetScores() (Scores, error) {
	resp, err := GetESPNScores()
	if err != nil {
		return Scores{}, fmt.Errorf("nflStats@GetScores: GetESPNScores, error: %s", err)
	}

	// Looking at fields in response
	fmt.Printf("main: GetESPNScores, resp: %+v", resp)
	// TODO: No Leagues
	v := reflect.ValueOf(resp)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
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

	espnNFLScoresResp := EspnNFLScores{}
	json.NewDecoder(resp.Body).Decode(espnNFLScoresResp)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: reading Get response body error: %s", err)
	}

	return EspnNFLScores{}, nil
}
