package nflstats

import (
	"fmt"
	"log"
	"net/http"
)

// GetScores fetches current NFL game scores
func GetScores() error {
	_, err := http.Get(nflURL(2021, 11, "REG"))
	return err
}

func nflURL(season, week int, seasonType string) string {
	base := "https://api.nfl.com/experience/v1/games"
	url := fmt.Sprintf("%s?season=%d&seasonType=%s&week=%d", base, season, seasonType, week)
	log.Printf("nflURL: %+v", url)
	return url
}
