package transformers

type Energon struct {
	Games []Game `json:"games"`
}

type Game struct {
	AwayScore int
	AwayTeam  string
	Digits    struct {
		Away1 int
		Away2 int
		Away3 int
		Away4 int
		Home1 int
		Home2 int
		Home3 int
		Home4 int
	}
	GameOver  bool
	HomeScore int
	HomeTeam  string
	ID        int
	Odds      struct {
		Details   string
		OverUnder float64
	}
	Quarter    int
	SeasonType int
	Time       string
	Week       int
	Year       int
}
