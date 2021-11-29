package main

import (
	"fmt"
)

func main() {
	scores, err := nflStats.GetScores()
	fmt.Println(scores)
	fmt.Println(err)
}
