package main

import (
	"fmt"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

func main() {
	games, err := nflStats.GetGames()
	fmt.Println(games)
	fmt.Println(err)
}
