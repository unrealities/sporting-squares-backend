package nflstats

// GetScores fetches current NFL game scores
func GetScores() (Scores, error) {
	URL := "http://static.nfl.com/liveupdate/scorestrip/scorestrip.json"
	resp, err := http.Get(URL)
	// TODO:
	// example response
	// {"ss":[["Sun","13:00:00","Final",,"MIA","26","BUF","56",,,"58408",,"REG17","2020"],["Sun","13:00:00","Final",,"BAL","38","CIN","3",,,"58411",,"REG17","2020"],["Sun","13:00:00","Final",,"PIT","22","CLE","24",,,"58412",,"REG17","2020"],["Sun","13:00:00","Final",,"MIN","37","DET","35",,,"58413",,"REG17","2020"],["Sun","13:00:00","Final",,"NYJ","14","NE","28",,,"58415",,"REG17","2020"],["Sun","13:00:00","Final",,"DAL","19","NYG","23",,,"58416",,"REG17","2020"],["Sun","13:00:00","Final",,"ATL","27","TB","44",,,"58418",,"REG17","2020"],["Sun","16:25:00","Final",,"NO","33","CAR","7",,,"58409",,"REG17","2020"],["Sun","16:25:00","Final",,"GB","35","CHI","16",,,"58410",,"REG17","2020"],["Sun","16:25:00","Final",,"LV","32","DEN","31",,,"58419",,"REG17","2020"],["Sun","16:25:00","Final",,"TEN","41","HOU","38",,,"58420",,"REG17","2020"],["Sun","16:25:00","Final",,"JAX","14","IND","28",,,"58421",,"REG17","2020"],["Sun","16:25:00","Final",,"LAC","38","KC","21",,,"58414",,"REG17","2020"],["Sun","16:25:00","Final",,"ARI","7","LA","18",,,"58422",,"REG17","2020"],["Sun","16:25:00","Final",,"SEA","26","SF","23",,,"58423",,"REG17","2020"],["Sun","20:20:00","Final",,"WAS","20","PHI","14",,,"58417",,"REG17","2020"]]}
}
