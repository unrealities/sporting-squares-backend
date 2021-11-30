package main

import (
	"fmt"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

func main() {
	scores, err := nflStats.GetScores()
	fmt.Println(scores)
	fmt.Println(err)
}
