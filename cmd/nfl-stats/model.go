package nflstats

// Game response example
// This is pulled from example-response.json
// That file was created during Thursday night football 11-11-2021 @ Halftime
// Lots of fluff in the JSON. Seems if there is a "detail" key, the game is in action
//
// "detail": {
// 	"id": "10160000-0586-39a4-241f-6d7470285c47",
// 	"gameClock": "15:00",
// 	"homePointsTotal": 6,
// 	"period": null,
// 	"phase": "HALFTIME",
// 	"possessionTeam": {
// 		"id": "10040325-2021-3010-b0e4-a25a32dbec2b",
// 		"abbreviation": "BAL"
// 	},
// 	"visitorPointsTotal": 3
// }
type Game struct {
	Detail struct {
		Id             string `json:"id"`
		GameClock      string `json:"gameClock"`
		HomeScore      string `json:"homePointsTotal"`
		Period         string `json:"period"`
		Phase          string `json:"phase"`
		PossessionTeam struct {
			Id           string `json:"id"`
			Abbreviation string `json:"abbreviation"`
		} `json:"possessionTeam"`
		VisitorScore string `json:"visitorPointsTotal"`
	} `json:"detail"`
}
