package nflStats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// TODO: persist to firebase cache storage
func GetGames() ([]Game, error) {
	games := []Game{}
	resp, err := GetESPNScores()
	if err != nil {
		return games, fmt.Errorf("nflStats@GetScores: GetESPNScores, error: %s", err)
	}

	for _, e := range resp.Events {
		g := Game{}
		// Tuesday after week the scores are still previous week
		// Wednesday new games are ready
		id, err := strconv.Atoi(e.ID)
		if err != nil {
			g.ID = -1
			continue
		}
		score, err := GetESPNGame(id)
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

						// TODO: This is verbose and duplicated
						firstQ, err := strconv.Atoi(t.Linescores[0].DisplayValue)
						if err != nil {
							g.Digits.Home1 = -1
						} else {
							g.Digits.Home1 = firstQ % 10
						}
						secondQ, err := strconv.Atoi(t.Linescores[1].DisplayValue)
						if err != nil {
							g.Digits.Home2 = -1
						} else {
							g.Digits.Home2 = (firstQ + secondQ) % 10
						}
						thirdQ, err := strconv.Atoi(t.Linescores[2].DisplayValue)
						if err != nil {
							g.Digits.Home3 = -1
						} else {
							g.Digits.Home3 = (firstQ + secondQ + thirdQ) % 10
						}
						fourthQ, err := strconv.Atoi(t.Linescores[3].DisplayValue)
						if err != nil {
							g.Digits.Home4 = -1
						} else {
							// TODO: index out of range. I assume the array is only length four if there is no overtime or the game hasn't started
							overtime, err := strconv.Atoi(t.Linescores[4].DisplayValue)
							if err != nil {
								g.Digits.Home4 = -1
							} else {
								g.Digits.Home4 = (firstQ + secondQ + thirdQ + fourthQ + overtime) % 10
							}
						}
					} else {
						g.AwayTeam = t.Team.Abbreviation
						g.AwayScore = score

						firstQ, err := strconv.Atoi(t.Linescores[0].DisplayValue)
						if err != nil {
							g.Digits.Away1 = -1
						} else {
							g.Digits.Away1 = firstQ % 10
						}
						secondQ, err := strconv.Atoi(t.Linescores[1].DisplayValue)
						if err != nil {
							g.Digits.Away2 = -1
						} else {
							g.Digits.Away2 = (firstQ + secondQ) % 10
						}
						thirdQ, err := strconv.Atoi(t.Linescores[2].DisplayValue)
						if err != nil {
							g.Digits.Away3 = -1
						} else {
							g.Digits.Away3 = (firstQ + secondQ + thirdQ) % 10
						}
						fourthQ, err := strconv.Atoi(t.Linescores[3].DisplayValue)
						if err != nil {
							g.Digits.Away4 = -1
						} else {
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
		g.ID = id                      // May be able to use ID to fetch individual box scores
		g.Quarter = e.Status.Period    // 0 = Game has not started. 4 = 4th or Game Over. 5 = Overtime or Game Over
		g.Time = e.Status.DisplayClock // 0 = Game has not started. 4 = 4th or Game Over. 5 = Overtime or Game Over
		g.GameOver = e.Status.Type.Completed
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

// GetESPNGame fetches and individual NFL game given a gameIdea
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
