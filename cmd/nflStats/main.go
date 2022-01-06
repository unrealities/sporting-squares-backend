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
func GetGames() ([]Game, error) {
	games := []Game{}

	resp, err := GetESPNGameByWeek(2, 17)
	if err != nil {
		return games, fmt.Errorf("nflStats@GetESPNGameByWeek: GetESPNGameByWeek, error: %s", err)
	}

	gameIDs := extractGameIDs(resp)

	for i, gameID := range gameIDs {
		g := Game{}

		score, err := GetESPNGame(gameIDs[i])
		if err != nil {
			g.Quarter = -1
			continue
		}

		for _, c := range score.Header.Competitions {
			if e.ID == c.ID {
				for _, t := range c.Competitors {
					score, err := strconv.Atoi(t.Score)
					if err != nil {
						score = -1
					}
					// To get last digit NUM % 10
					if t.HomeAway == "home" {
						g.HomeTeam = t.Team.Abbreviation
						g.HomeScore = score

						if len(t.Linescores) == 0 {
							g.Digits.Home1 = -1
							g.Digits.Home2 = -1
							g.Digits.Home3 = -1
							g.Digits.Home4 = -1
							continue
						}
						// TODO: This is verbose and duplicated
						firstQ, err := strconv.Atoi(t.Linescores[0].DisplayValue)
						if err != nil {
							g.Digits.Home1 = -1
						} else {
							g.Digits.Home1 = firstQ % 10
						}
						if len(t.Linescores) == 1 {
							g.Digits.Home2 = -1
							g.Digits.Home3 = -1
							g.Digits.Home4 = -1
							continue
						}
						secondQ, err := strconv.Atoi(t.Linescores[1].DisplayValue)
						if err != nil {
							g.Digits.Home2 = -1
						} else {
							g.Digits.Home2 = (firstQ + secondQ) % 10
						}
						if len(t.Linescores) == 2 {
							g.Digits.Home3 = -1
							g.Digits.Home4 = -1
							continue
						}
						thirdQ, err := strconv.Atoi(t.Linescores[2].DisplayValue)
						if err != nil {
							g.Digits.Home3 = -1
						} else {
							g.Digits.Home3 = (firstQ + secondQ + thirdQ) % 10
						}
						if len(t.Linescores) == 3 {
							g.Digits.Home4 = -1
							continue
						}
						fourthQ, err := strconv.Atoi(t.Linescores[3].DisplayValue)
						if err != nil {
							g.Digits.Home4 = -1
						} else {
							g.Digits.Home4 = (firstQ + secondQ + thirdQ + fourthQ) % 10
							if len(t.Linescores) > 4 {
								overtime, err := strconv.Atoi(t.Linescores[4].DisplayValue)
								if err != nil {
									g.Digits.Home4 = -1
								} else {
									g.Digits.Home4 = (firstQ + secondQ + thirdQ + fourthQ + overtime) % 10
								}
							}
						}
					} else {
						g.AwayTeam = t.Team.Abbreviation
						g.AwayScore = score

						if len(t.Linescores) == 0 {
							g.Digits.Away1 = -1
							g.Digits.Away2 = -1
							g.Digits.Away3 = -1
							g.Digits.Away4 = -1
							continue
						}
						firstQ, err := strconv.Atoi(t.Linescores[0].DisplayValue)
						if err != nil {
							g.Digits.Away1 = -1
						} else {
							g.Digits.Away1 = firstQ % 10
						}
						if len(t.Linescores) == 1 {
							g.Digits.Away2 = -1
							g.Digits.Away3 = -1
							g.Digits.Away4 = -1
							continue
						}
						secondQ, err := strconv.Atoi(t.Linescores[1].DisplayValue)
						if err != nil {
							g.Digits.Away2 = -1
						} else {
							g.Digits.Away2 = (firstQ + secondQ) % 10
						}
						if len(t.Linescores) == 2 {
							g.Digits.Away3 = -1
							g.Digits.Away4 = -1
							continue
						}
						thirdQ, err := strconv.Atoi(t.Linescores[2].DisplayValue)
						if err != nil {
							g.Digits.Away3 = -1
						} else {
							g.Digits.Away3 = (firstQ + secondQ + thirdQ) % 10
						}
						if len(t.Linescores) == 3 {
							g.Digits.Away4 = -1
							continue
						}
						fourthQ, err := strconv.Atoi(t.Linescores[3].DisplayValue)
						if err != nil {
							g.Digits.Away4 = -1
						} else {
							g.Digits.Away4 = (firstQ + secondQ + thirdQ + fourthQ) % 10
							if len(t.Linescores) > 4 {
								overtime, err := strconv.Atoi(t.Linescores[4].DisplayValue)
								if err != nil {
									g.Digits.Away4 = -1
								} else {
									g.Digits.Away4 = (firstQ + secondQ + thirdQ + fourthQ + overtime) % 10
								}
							}
						}
					}
				}
			}
		}

		g.ID = gameID
		/*  TODO: Get this from individual game
		g.Quarter = e.Status.Period    // 0 = Game has not started. 4 = 4th or Game Over. 5 = Overtime or Game Over
		g.Time = e.Status.DisplayClock // 0 = Game has not started. 4 = 4th or Game Over. 5 = Overtime or Game Over
		g.GameOver = e.Status.Type.Completed
		*/
		for _, c := range e.Competitions {
			if len(c.Odds) > 0 {
				g.Odds.Details = c.Odds[0].Details     // ex. ARI -2.5 (need to separate out team and spread)
				g.Odds.OverUnder = c.Odds[0].OverUnder // ex. 51.5
			}
			games = append(games, g)
		}
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
func extractGameIDs(games ESPNNFLGamesByWeek) ([]int) {
	gameIDs := []int{}
	for _, link := range games.Items {
		gameID, _ := regexp.MatchString("[^events\/](?:[0-9].*[0-9])", link)
		gameIDs = append(gameIDs, gameID)
	}
	return gameIDs
}
