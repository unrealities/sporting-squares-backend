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

// TODO: This can get the games for a given week: http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2021/types/2/weeks/17/events?lang=en&region=us
/*
{"$meta":{"parameters":{"week":["17"],"season":["2021"],"seasontypes":["2"]}},"count":16,"pageIndex":1,"pageSize":25,"pageCount":1,"items":[{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326569?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326570?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326571?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326577?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326573?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326574?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326576?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326578?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326581?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326579?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326580?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326572?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326575?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326582?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326583?lang=en&region=us"},{"$ref":"http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401326584?lang=en&region=us"}]}
*/
func GetESPNGameByWeek(week int) (ESPNNFLGamesByWeek, error) {
	URL := fmt.Printf("http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2021/types/2/weeks/%d/events?lang=en&region=us", week)
	resp, err := http.Get(URL)
	if err != nil {
		return ESPNNFLGamesByWeek{}, fmt.Errorf("nflStats#GetESPNGameByWeek: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
}
