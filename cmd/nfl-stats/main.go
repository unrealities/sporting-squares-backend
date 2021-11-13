package nflstats

import (
	"fmt"
	"net/http"
)

// GetScores fetches current NFL game scores
func GetScores() error {
	_, err := http.Get(nflURL(2021, 11))
	return err
}

func nflURL(season, week int) string {
	base := "https://api.nfl.com/experience/v1/games"
	return fmt.Sprintf("%s?season=%d&seasonType=REG&week=%d", base, season, week)
}
