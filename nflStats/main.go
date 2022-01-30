package nflStats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// TODO: persist to firebase cache storage
func GetGames() ([]EspnNFLGame, error) {
	games := []EspnNFLGame{}

	week := 3
	seasonType := 3
	// TODO: what am I passing to the transformer?
	resp, err := GetESPNGameByWeek(seasonType, week)
	if err != nil {
		return games, fmt.Errorf("nflStats@GetESPNGameByWeek: GetESPNGameByWeek, error: %s", err)
	}

	return games, nil
}

// GetESPNGame fetches an individual NFL game given an ESPN gameID
func GetESPNGame(gameID int) (EspnNFLGame, error) {
	URL := "https://site.api.espn.com/apis/site/v2/sports/football/nfl/summary?event=" + strconv.Itoa(gameID)
	resp, err := http.Get(URL)
	if err != nil {
		return EspnNFLGame{}, fmt.Errorf("nflStats#GetESPNGame: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	espnNFLGameResp := EspnNFLGame{}
	json.Unmarshal([]byte(body), &espnNFLGameResp)
	if err != nil {
		return EspnNFLGame{}, fmt.Errorf("nflStats#GetESPNGame: reading Get response body error: %s", err)
	}

	return espnNFLGameResp, nil
}

// GetESPNGameByWeek fetches games for a given week and seasonType
// seasonType 1 = preseason
// seasonType 2 = regular season
// seasonType 3 = playoffs
// week is relative to the seasonType. i.e. {seasonType: 1, week: 1} = The Hall of Fame Game
func GetESPNGameByWeek(seasonType, week int) (ESPNNFLGamesByWeek, error) {
	URL := fmt.Sprintf("http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2021/types/%d/weeks/%d/events?lang=en&region=us", seasonType, week)
	resp, err := http.Get(URL)
	if err != nil {
		return ESPNNFLGamesByWeek{}, fmt.Errorf("nflStats#GetESPNGameByWeek: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	espnNFLGamesByWeekResp := ESPNNFLGamesByWeek{}
	json.Unmarshal([]byte(body), &espnNFLGamesByWeekResp)
	if err != nil {
		return ESPNNFLGamesByWeek{}, fmt.Errorf("nflStats#GetESPNGameByWeek: reading Get response body error: %s", err)
	}

	return espnNFLGamesByWeekResp, nil
}

// extractGameIDs parses gameIDs out of ESPNNFLGamesByWeek
func extractGameIDs(games ESPNNFLGamesByWeek) []int {
	gameIDs := []int{}
	r, err := regexp.Compile("[^events/](?:[0-9].*[0-9])")
	if err != nil {
		return gameIDs
	}
	for _, link := range games.Items {
		strGameID := r.FindString(link.Ref)
		gameID, err := strconv.Atoi(strGameID)
		if err != nil {
			gameID = -1
		}
		gameIDs = append(gameIDs, gameID)
	}
	return gameIDs
}
