package nflStats

import "time"

type EspnNFLScores struct {
	Leagues []struct {
		ID           string `json:"id"`
		UID          string `json:"uid"`
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		Slug         string `json:"slug"`
		Season       struct {
			Year      int    `json:"year"`
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
			Type      struct {
				ID           string `json:"id"`
				Type         int    `json:"type"`
				Name         string `json:"name"`
				Abbreviation string `json:"abbreviation"`
			} `json:"type"`
		} `json:"season"`
		CalendarType        string `json:"calendarType"`
		CalendarIsWhitelist bool   `json:"calendarIsWhitelist"`
		CalendarStartDate   string `json:"calendarStartDate"`
		CalendarEndDate     string `json:"calendarEndDate"`
		Calendar            []struct {
			Label     string `json:"label"`
			Value     string `json:"value"`
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
			Entries   []struct {
				Label          string `json:"label"`
				AlternateLabel string `json:"alternateLabel"`
				Detail         string `json:"detail"`
				Value          string `json:"value"`
				StartDate      string `json:"startDate"`
				EndDate        string `json:"endDate"`
			} `json:"entries"`
		} `json:"calendar"`
	} `json:"leagues"`
	Season struct {
		Type int `json:"type"`
		Year int `json:"year"`
	} `json:"season"`
	Week struct {
		Number     int `json:"number"`
		TeamsOnBye []struct {
			ID               string `json:"id"`
			UID              string `json:"uid"`
			Location         string `json:"location"`
			Name             string `json:"name"`
			Abbreviation     string `json:"abbreviation"`
			DisplayName      string `json:"displayName"`
			ShortDisplayName string `json:"shortDisplayName"`
			IsActive         bool   `json:"isActive"`
			Links            []struct {
				Rel        []string `json:"rel"`
				Href       string   `json:"href"`
				Text       string   `json:"text"`
				IsExternal bool     `json:"isExternal"`
				IsPremium  bool     `json:"isPremium"`
			} `json:"links"`
			Logo string `json:"logo"`
		} `json:"teamsOnBye"`
	} `json:"week"`
	Events []struct {
		ID        string `json:"id"`
		UID       string `json:"uid"`
		Date      string `json:"date"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Season    struct {
			Year int    `json:"year"`
			Type int    `json:"type"`
			Slug string `json:"slug"`
		} `json:"season"`
		Competitions []struct {
			ID         string `json:"id"`
			UID        string `json:"uid"`
			Date       string `json:"date"`
			Attendance int    `json:"attendance"`
			Type       struct {
				ID           string `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"type"`
			TimeValid             bool `json:"timeValid"`
			NeutralSite           bool `json:"neutralSite"`
			ConferenceCompetition bool `json:"conferenceCompetition"`
			Recent                bool `json:"recent"`
			Venue                 struct {
				ID       string `json:"id"`
				FullName string `json:"fullName"`
				Address  struct {
					City  string `json:"city"`
					State string `json:"state"`
				} `json:"address"`
				Capacity int  `json:"capacity"`
				Indoor   bool `json:"indoor"`
			} `json:"venue"`
			Competitors []struct {
				ID       string `json:"id"`
				UID      string `json:"uid"`
				Type     string `json:"type"`
				Order    int    `json:"order"`
				HomeAway string `json:"homeAway"`
				Team     struct {
					ID               string `json:"id"`
					UID              string `json:"uid"`
					Location         string `json:"location"`
					Name             string `json:"name"`
					Abbreviation     string `json:"abbreviation"`
					DisplayName      string `json:"displayName"`
					ShortDisplayName string `json:"shortDisplayName"`
					Color            string `json:"color"`
					AlternateColor   string `json:"alternateColor"`
					IsActive         bool   `json:"isActive"`
					Venue            struct {
						ID string `json:"id"`
					} `json:"venue"`
					Links []struct {
						Rel        []string `json:"rel"`
						Href       string   `json:"href"`
						Text       string   `json:"text"`
						IsExternal bool     `json:"isExternal"`
						IsPremium  bool     `json:"isPremium"`
					} `json:"links"`
					Logo string `json:"logo"`
				} `json:"team"`
				Score      string        `json:"score"`
				Statistics []interface{} `json:"statistics"`
				Records    []struct {
					Name         string `json:"name"`
					Abbreviation string `json:"abbreviation,omitempty"`
					Type         string `json:"type"`
					Summary      string `json:"summary"`
				} `json:"records"`
				Leaders []struct {
					Name             string `json:"name"`
					DisplayName      string `json:"displayName"`
					ShortDisplayName string `json:"shortDisplayName"`
					Abbreviation     string `json:"abbreviation"`
					Leaders          []struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
						Athlete      struct {
							ID          string `json:"id"`
							FullName    string `json:"fullName"`
							DisplayName string `json:"displayName"`
							ShortName   string `json:"shortName"`
							Links       []struct {
								Rel  []string `json:"rel"`
								Href string   `json:"href"`
							} `json:"links"`
							Headshot string `json:"headshot"`
							Jersey   string `json:"jersey"`
							Position struct {
								Abbreviation string `json:"abbreviation"`
							} `json:"position"`
							Team struct {
								ID string `json:"id"`
							} `json:"team"`
							Active bool `json:"active"`
						} `json:"athlete"`
						Team struct {
							ID string `json:"id"`
						} `json:"team"`
					} `json:"leaders"`
				} `json:"leaders"`
			} `json:"competitors"`
			Notes  []interface{} `json:"notes"`
			Status struct {
				Clock        float64 `json:"clock"`
				DisplayClock string  `json:"displayClock"`
				Period       int     `json:"period"`
				Type         struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					State       string `json:"state"`
					Completed   bool   `json:"completed"`
					Description string `json:"description"`
					Detail      string `json:"detail"`
					ShortDetail string `json:"shortDetail"`
				} `json:"type"`
			} `json:"status"`
			Broadcasts []struct {
				Market string   `json:"market"`
				Names  []string `json:"names"`
			} `json:"broadcasts"`
			Leaders []struct {
				Name             string `json:"name"`
				DisplayName      string `json:"displayName"`
				ShortDisplayName string `json:"shortDisplayName"`
				Abbreviation     string `json:"abbreviation"`
				Leaders          []struct {
					DisplayValue string  `json:"displayValue"`
					Value        float64 `json:"value"`
					Athlete      struct {
						ID          string `json:"id"`
						FullName    string `json:"fullName"`
						DisplayName string `json:"displayName"`
						ShortName   string `json:"shortName"`
						Links       []struct {
							Rel  []string `json:"rel"`
							Href string   `json:"href"`
						} `json:"links"`
						Headshot string `json:"headshot"`
						Jersey   string `json:"jersey"`
						Position struct {
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
						Team struct {
							ID string `json:"id"`
						} `json:"team"`
						Active bool `json:"active"`
					} `json:"athlete"`
					Team struct {
						ID string `json:"id"`
					} `json:"team"`
				} `json:"leaders"`
			} `json:"leaders"`
			Tickets []struct {
				Summary         string `json:"summary"`
				NumberAvailable int    `json:"numberAvailable"`
				Links           []struct {
					Href string `json:"href"`
				} `json:"links"`
			} `json:"tickets"`
			StartDate     string `json:"startDate"`
			GeoBroadcasts []struct {
				Type struct {
					ID        string `json:"id"`
					ShortName string `json:"shortName"`
				} `json:"type"`
				Market struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"market"`
				Media struct {
					ShortName string `json:"shortName"`
				} `json:"media"`
				Lang   string `json:"lang"`
				Region string `json:"region"`
			} `json:"geoBroadcasts"`
			Odds []struct {
				Provider struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Priority int    `json:"priority"`
				} `json:"provider"`
				Details   string  `json:"details"`
				OverUnder float64 `json:"overUnder"`
			} `json:"odds"`
		} `json:"competitions"`
		Links []struct {
			Language   string   `json:"language"`
			Rel        []string `json:"rel"`
			Href       string   `json:"href"`
			Text       string   `json:"text"`
			ShortText  string   `json:"shortText"`
			IsExternal bool     `json:"isExternal"`
			IsPremium  bool     `json:"isPremium"`
		} `json:"links"`
		Weather struct {
			DisplayValue    string `json:"displayValue"`
			Temperature     int    `json:"temperature"`
			HighTemperature int    `json:"highTemperature"`
			ConditionID     string `json:"conditionId"`
			Link            struct {
				Language   string   `json:"language"`
				Rel        []string `json:"rel"`
				Href       string   `json:"href"`
				Text       string   `json:"text"`
				ShortText  string   `json:"shortText"`
				IsExternal bool     `json:"isExternal"`
				IsPremium  bool     `json:"isPremium"`
			} `json:"link"`
		} `json:"weather,omitempty"`
		Status struct {
			Clock        float64 `json:"clock"`
			DisplayClock string  `json:"displayClock"`
			Period       int     `json:"period"`
			Type         struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				State       string `json:"state"`
				Completed   bool   `json:"completed"`
				Description string `json:"description"`
				Detail      string `json:"detail"`
				ShortDetail string `json:"shortDetail"`
			} `json:"type"`
		} `json:"status"`
	} `json:"events"`
}

type EspnNFLGame struct {
	Boxscore struct {
		Teams []struct {
			Team struct {
				ID               string `json:"id"`
				UID              string `json:"uid"`
				Slug             string `json:"slug"`
				Location         string `json:"location"`
				Name             string `json:"name"`
				Abbreviation     string `json:"abbreviation"`
				DisplayName      string `json:"displayName"`
				ShortDisplayName string `json:"shortDisplayName"`
				Color            string `json:"color"`
				AlternateColor   string `json:"alternateColor"`
				Logo             string `json:"logo"`
			} `json:"team"`
			Statistics []struct {
				Name         string `json:"name"`
				DisplayValue string `json:"displayValue"`
				Label        string `json:"label"`
			} `json:"statistics"`
		} `json:"teams"`
	} `json:"boxscore"`
	GameInfo struct {
		Venue struct {
			ID       string `json:"id"`
			FullName string `json:"fullName"`
			Address  struct {
				City    string `json:"city"`
				State   string `json:"state"`
				ZipCode string `json:"zipCode"`
			} `json:"address"`
			Capacity int  `json:"capacity"`
			Grass    bool `json:"grass"`
			Images   []struct {
				Href   string   `json:"href"`
				Width  int      `json:"width"`
				Height int      `json:"height"`
				Alt    string   `json:"alt"`
				Rel    []string `json:"rel"`
			} `json:"images"`
		} `json:"venue"`
		Weather struct {
			Temperature     int    `json:"temperature"`
			HighTemperature int    `json:"highTemperature"`
			LowTemperature  int    `json:"lowTemperature"`
			ConditionID     string `json:"conditionId"`
			Gust            int    `json:"gust"`
			Precipitation   int    `json:"precipitation"`
			Link            struct {
				Language   string   `json:"language"`
				Rel        []string `json:"rel"`
				Href       string   `json:"href"`
				Text       string   `json:"text"`
				ShortText  string   `json:"shortText"`
				IsExternal bool     `json:"isExternal"`
				IsPremium  bool     `json:"isPremium"`
			} `json:"link"`
		} `json:"weather"`
	} `json:"gameInfo"`
	LastFiveGames []struct {
		Team struct {
			ID           string `json:"id"`
			UID          string `json:"uid"`
			DisplayName  string `json:"displayName"`
			Abbreviation string `json:"abbreviation"`
			Links        []struct {
				Href string `json:"href"`
				Text string `json:"text"`
			} `json:"links"`
			Logo string `json:"logo"`
		} `json:"team"`
		Events []struct {
			ID    string `json:"id"`
			Links []struct {
				Href string `json:"href"`
				Text string `json:"text"`
			} `json:"links"`
			Week               int    `json:"week"`
			AtVs               string `json:"atVs"`
			GameDate           string `json:"gameDate"`
			Score              string `json:"score"`
			HomeTeamID         string `json:"homeTeamId"`
			AwayTeamID         string `json:"awayTeamId"`
			HomeTeamScore      string `json:"homeTeamScore"`
			AwayTeamScore      string `json:"awayTeamScore"`
			HomeAggregateScore string `json:"homeAggregateScore"`
			AwayAggregateScore string `json:"awayAggregateScore"`
			HomeShootoutScore  string `json:"homeShootoutScore"`
			AwayShootoutScore  string `json:"awayShootoutScore"`
			GameResult         string `json:"gameResult"`
			Opponent           struct {
				ID           string `json:"id"`
				UID          string `json:"uid"`
				DisplayName  string `json:"displayName"`
				Abbreviation string `json:"abbreviation"`
				Links        []struct {
					Href string `json:"href"`
					Text string `json:"text"`
				} `json:"links"`
				Logo string `json:"logo"`
			} `json:"opponent"`
			OpponentLogo       string `json:"opponentLogo"`
			LeagueName         string `json:"leagueName"`
			LeagueAbbreviation string `json:"leagueAbbreviation"`
		} `json:"events"`
	} `json:"lastFiveGames"`
	Leaders []struct {
		Team struct {
			ID           string `json:"id"`
			UID          string `json:"uid"`
			DisplayName  string `json:"displayName"`
			Abbreviation string `json:"abbreviation"`
			Links        []struct {
				Href string `json:"href"`
				Text string `json:"text"`
			} `json:"links"`
			Logo string `json:"logo"`
		} `json:"team"`
		Leaders []struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			Leaders     []struct {
				DisplayValue string `json:"displayValue"`
				Athlete      struct {
					ID          string `json:"id"`
					UID         string `json:"uid"`
					GUID        string `json:"guid"`
					LastName    string `json:"lastName"`
					FullName    string `json:"fullName"`
					DisplayName string `json:"displayName"`
					ShortName   string `json:"shortName"`
					Links       []struct {
						Rel  []string `json:"rel"`
						Href string   `json:"href"`
						Text string   `json:"text"`
					} `json:"links"`
					Headshot struct {
						Href string `json:"href"`
						Alt  string `json:"alt"`
					} `json:"headshot"`
					Jersey   string `json:"jersey"`
					Position struct {
						Abbreviation string `json:"abbreviation"`
					} `json:"position"`
				} `json:"athlete"`
			} `json:"leaders"`
		} `json:"leaders"`
	} `json:"leaders"`
	Injuries []struct {
		Team struct {
			ID           string `json:"id"`
			UID          string `json:"uid"`
			DisplayName  string `json:"displayName"`
			Abbreviation string `json:"abbreviation"`
			Links        []struct {
				Href string `json:"href"`
				Text string `json:"text"`
			} `json:"links"`
			Logo string `json:"logo"`
		} `json:"team"`
		Injuries []struct {
			Status  string `json:"status"`
			Date    string `json:"date"`
			Athlete struct {
				ID          string `json:"id"`
				UID         string `json:"uid"`
				GUID        string `json:"guid"`
				LastName    string `json:"lastName"`
				FullName    string `json:"fullName"`
				DisplayName string `json:"displayName"`
				ShortName   string `json:"shortName"`
				Links       []struct {
					Rel  []string `json:"rel"`
					Href string   `json:"href"`
					Text string   `json:"text"`
				} `json:"links"`
				Headshot struct {
					Href string `json:"href"`
					Alt  string `json:"alt"`
				} `json:"headshot"`
				Jersey   string `json:"jersey"`
				Position struct {
					Name         string `json:"name"`
					DisplayName  string `json:"displayName"`
					Abbreviation string `json:"abbreviation"`
				} `json:"position"`
			} `json:"athlete"`
			Type struct {
				ID           string `json:"id"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Abbreviation string `json:"abbreviation"`
			} `json:"type"`
		} `json:"injuries"`
	} `json:"injuries"`
	Broadcasts []struct {
		Type struct {
			ID        string `json:"id"`
			ShortName string `json:"shortName"`
			LongName  string `json:"longName"`
			Slug      string `json:"slug"`
		} `json:"type"`
		Station string `json:"station"`
		Market  struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"market"`
		Media struct {
			CallLetters string `json:"callLetters"`
			Name        string `json:"name"`
			ShortName   string `json:"shortName"`
		} `json:"media"`
		Lang   string `json:"lang"`
		Region string `json:"region"`
	} `json:"broadcasts"`
	Predictor struct {
		Header   string `json:"header"`
		HomeTeam struct {
			ID             string `json:"id"`
			GameProjection string `json:"gameProjection"`
			TeamChanceLoss string `json:"teamChanceLoss"`
			TeamChanceTie  string `json:"teamChanceTie"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID             string `json:"id"`
			GameProjection string `json:"gameProjection"`
			TeamChanceLoss string `json:"teamChanceLoss"`
			TeamChanceTie  string `json:"teamChanceTie"`
		} `json:"awayTeam"`
	} `json:"predictor"`
	Pickcenter []struct {
		Provider struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Priority int    `json:"priority"`
		} `json:"provider"`
		Details      string  `json:"details"`
		OverUnder    float64 `json:"overUnder"`
		Spread       float64 `json:"spread"`
		AwayTeamOdds struct {
			Favorite   bool   `json:"favorite"`
			Underdog   bool   `json:"underdog"`
			MoneyLine  int    `json:"moneyLine"`
			SpreadOdds int    `json:"spreadOdds"`
			TeamID     string `json:"teamId"`
		} `json:"awayTeamOdds,omitempty"`
		HomeTeamOdds struct {
			Favorite   bool   `json:"favorite"`
			Underdog   bool   `json:"underdog"`
			MoneyLine  int    `json:"moneyLine"`
			SpreadOdds int    `json:"spreadOdds"`
			TeamID     string `json:"teamId"`
		} `json:"homeTeamOdds,omitempty"`
		Links []interface{} `json:"links"`
	} `json:"pickcenter"`
	AgainstTheSpread []struct {
		Team struct {
			ID           string `json:"id"`
			UID          string `json:"uid"`
			DisplayName  string `json:"displayName"`
			Abbreviation string `json:"abbreviation"`
			Links        []struct {
				Href string `json:"href"`
				Text string `json:"text"`
			} `json:"links"`
			Logo string `json:"logo"`
		} `json:"team"`
		Records []interface{} `json:"records"`
	} `json:"againstTheSpread"`
	Odds []struct {
		Provider struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Priority int    `json:"priority"`
		} `json:"provider"`
		BettingOdds struct {
			HomeTeam struct {
				Ref string `json:"$ref"`
			} `json:"homeTeam"`
			AwayTeam struct {
				Ref string `json:"$ref"`
			} `json:"awayTeam"`
			TeamOdds struct {
				PreMatchMoneyLineAway struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchMoneyLineAway"`
				PreMatchMoneyLineHome struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchMoneyLineHome"`
				PreMatchSpreadHandicapAway struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchSpreadHandicapAway"`
				PreMatchSpreadHome struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchSpreadHome"`
				PreMatchWinningMarginHome struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchWinningMarginHome"`
				PreMatchTotalOver struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchTotalOver"`
				PreMatchWinningMarginOther struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchWinningMarginOther"`
				PreMatchSpreadAway struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchSpreadAway"`
				PreMatchWinningMarginAway struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchWinningMarginAway"`
				PreMatchTotalUnder struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchTotalUnder"`
				PreMatchTotalHandicap struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchTotalHandicap"`
				PreMatchSpreadHandicapHome struct {
					OddID      string `json:"oddId"`
					Value      string `json:"value"`
					BetSlipURL string `json:"betSlipUrl"`
				} `json:"preMatchSpreadHandicapHome"`
			} `json:"teamOdds"`
		} `json:"bettingOdds"`
	} `json:"odds"`
	Header struct {
		ID     string `json:"id"`
		UID    string `json:"uid"`
		Season struct {
			Year int `json:"year"`
			Type int `json:"type"`
		} `json:"season"`
		TimeValid    bool `json:"timeValid"`
		Competitions []struct {
			ID                    string `json:"id"`
			UID                   string `json:"uid"`
			Date                  string `json:"date"`
			NeutralSite           bool   `json:"neutralSite"`
			ConferenceCompetition bool   `json:"conferenceCompetition"`
			BoxscoreAvailable     bool   `json:"boxscoreAvailable"`
			CommentaryAvailable   bool   `json:"commentaryAvailable"`
			LiveAvailable         bool   `json:"liveAvailable"`
			OnWatchESPN           bool   `json:"onWatchESPN"`
			Recent                bool   `json:"recent"`
			BoxscoreSource        string `json:"boxscoreSource"`
			PlayByPlaySource      string `json:"playByPlaySource"`
			Competitors           []struct {
				ID       string `json:"id"`
				UID      string `json:"uid"`
				Order    int    `json:"order"`
				HomeAway string `json:"homeAway"`
				Team     struct {
					ID             string `json:"id"`
					UID            string `json:"uid"`
					Location       string `json:"location"`
					Name           string `json:"name"`
					Nickname       string `json:"nickname"`
					Abbreviation   string `json:"abbreviation"`
					DisplayName    string `json:"displayName"`
					Color          string `json:"color"`
					AlternateColor string `json:"alternateColor"`
					Logos          []struct {
						Href        string   `json:"href"`
						Width       int      `json:"width"`
						Height      int      `json:"height"`
						Alt         string   `json:"alt"`
						Rel         []string `json:"rel"`
						LastUpdated string   `json:"lastUpdated"`
					} `json:"logos"`
					Links []struct {
						Rel  []string `json:"rel"`
						Href string   `json:"href"`
						Text string   `json:"text"`
					} `json:"links"`
				} `json:"team"`
				Record []struct {
					Type         string `json:"type"`
					Summary      string `json:"summary"`
					DisplayValue string `json:"displayValue"`
				} `json:"record"`
				Possession bool `json:"possession"`
			} `json:"competitors"`
			Status struct {
				Type struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					State       string `json:"state"`
					Completed   bool   `json:"completed"`
					Description string `json:"description"`
					Detail      string `json:"detail"`
					ShortDetail string `json:"shortDetail"`
				} `json:"type"`
			} `json:"status"`
			Broadcasts []struct {
				Type struct {
					ID        string `json:"id"`
					ShortName string `json:"shortName"`
				} `json:"type"`
				Market struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"market"`
				Media struct {
					ShortName string `json:"shortName"`
				} `json:"media"`
				Lang   string `json:"lang"`
				Region string `json:"region"`
			} `json:"broadcasts"`
		} `json:"competitions"`
		Links []struct {
			Rel        []string `json:"rel"`
			Href       string   `json:"href"`
			Text       string   `json:"text"`
			ShortText  string   `json:"shortText"`
			IsExternal bool     `json:"isExternal"`
			IsPremium  bool     `json:"isPremium"`
		} `json:"links"`
		Week   int `json:"week"`
		League struct {
			ID           string `json:"id"`
			UID          string `json:"uid"`
			Name         string `json:"name"`
			Abbreviation string `json:"abbreviation"`
			Slug         string `json:"slug"`
			IsTournament bool   `json:"isTournament"`
			Links        []struct {
				Rel  []string `json:"rel"`
				Href string   `json:"href"`
				Text string   `json:"text"`
			} `json:"links"`
		} `json:"league"`
	} `json:"header"`
	News struct {
		Header string `json:"header"`
		Link   struct {
			Language   string   `json:"language"`
			Rel        []string `json:"rel"`
			Href       string   `json:"href"`
			Text       string   `json:"text"`
			ShortText  string   `json:"shortText"`
			IsExternal bool     `json:"isExternal"`
			IsPremium  bool     `json:"isPremium"`
		} `json:"link"`
		Articles []struct {
			Images []struct {
				Name   string `json:"name"`
				Width  int    `json:"width"`
				ID     int    `json:"id"`
				Credit string `json:"credit"`
				Type   string `json:"type"`
				URL    string `json:"url"`
				Height int    `json:"height"`
			} `json:"images"`
			Description  string    `json:"description"`
			Published    time.Time `json:"published"`
			Type         string    `json:"type"`
			Premium      bool      `json:"premium"`
			LastModified time.Time `json:"lastModified"`
			Categories   []struct {
				ID          int       `json:"id,omitempty"`
				Description string    `json:"description,omitempty"`
				Type        string    `json:"type"`
				SportID     int       `json:"sportId,omitempty"`
				TopicID     int       `json:"topicId,omitempty"`
				CreateDate  time.Time `json:"createDate"`
				LeagueID    int       `json:"leagueId,omitempty"`
				League      struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Links       struct {
						API struct {
							Leagues struct {
								Href string `json:"href"`
							} `json:"leagues"`
						} `json:"api"`
						Web struct {
							Leagues struct {
								Href string `json:"href"`
							} `json:"leagues"`
						} `json:"web"`
						Mobile struct {
							Leagues struct {
								Href string `json:"href"`
							} `json:"leagues"`
						} `json:"mobile"`
					} `json:"links"`
				} `json:"league,omitempty"`
				UID    string `json:"uid,omitempty"`
				TeamID int    `json:"teamId,omitempty"`
				Team   struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Links       struct {
						API struct {
							Teams struct {
								Href string `json:"href"`
							} `json:"teams"`
						} `json:"api"`
						Web struct {
							Teams struct {
								Href string `json:"href"`
							} `json:"teams"`
						} `json:"web"`
						Mobile struct {
							Teams struct {
								Href string `json:"href"`
							} `json:"teams"`
						} `json:"mobile"`
					} `json:"links"`
				} `json:"team,omitempty"`
				GUID string `json:"guid,omitempty"`
			} `json:"categories"`
			Headline string `json:"headline"`
			Byline   string `json:"byline,omitempty"`
			Links    struct {
				API struct {
					News struct {
						Href string `json:"href"`
					} `json:"news"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"api"`
				Web struct {
					Href  string `json:"href"`
					Short struct {
						Href string `json:"href"`
					} `json:"short"`
				} `json:"web"`
			} `json:"links,omitempty"`
		} `json:"articles"`
	} `json:"news"`
	Article struct {
		Keywords    []interface{} `json:"keywords"`
		Description string        `json:"description"`
		Source      string        `json:"source"`
		Video       []struct {
			Source      string `json:"source"`
			ID          int    `json:"id"`
			Headline    string `json:"headline"`
			Caption     string `json:"caption"`
			Description string `json:"description"`
			Premium     bool   `json:"premium"`
			Ad          struct {
				Sport  string `json:"sport"`
				Bundle string `json:"bundle"`
			} `json:"ad"`
			Tracking struct {
				SportName    string `json:"sportName"`
				LeagueName   string `json:"leagueName"`
				CoverageType string `json:"coverageType"`
				TrackingName string `json:"trackingName"`
				TrackingID   string `json:"trackingId"`
			} `json:"tracking"`
			CerebroID           string    `json:"cerebroId"`
			LastModified        time.Time `json:"lastModified"`
			OriginalPublishDate time.Time `json:"originalPublishDate"`
			TimeRestrictions    struct {
				EmbargoDate    time.Time `json:"embargoDate"`
				ExpirationDate time.Time `json:"expirationDate"`
			} `json:"timeRestrictions"`
			DeviceRestrictions struct {
				Type    string   `json:"type"`
				Devices []string `json:"devices"`
			} `json:"deviceRestrictions"`
			Syndicatable bool `json:"syndicatable"`
			Duration     int  `json:"duration"`
			Categories   []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Type        string `json:"type"`
				SportID     int    `json:"sportId"`
				LeagueID    int    `json:"leagueId,omitempty"`
				League      struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Links       struct {
						API struct {
							Leagues struct {
								Href string `json:"href"`
							} `json:"leagues"`
						} `json:"api"`
						Web struct {
							Leagues struct {
								Href string `json:"href"`
							} `json:"leagues"`
						} `json:"web"`
						Mobile struct {
							Leagues struct {
								Href string `json:"href"`
							} `json:"leagues"`
						} `json:"mobile"`
					} `json:"links"`
				} `json:"league,omitempty"`
				UID    string `json:"uid"`
				TeamID int    `json:"teamId,omitempty"`
				Team   struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Links       struct {
						API struct {
							Teams struct {
								Href string `json:"href"`
							} `json:"teams"`
						} `json:"api"`
						Web struct {
							Teams struct {
								Href string `json:"href"`
							} `json:"teams"`
						} `json:"web"`
						Mobile struct {
							Teams struct {
								Href string `json:"href"`
							} `json:"teams"`
						} `json:"mobile"`
					} `json:"links"`
				} `json:"team,omitempty"`
			} `json:"categories"`
			Keywords     []interface{} `json:"keywords"`
			PosterImages struct {
				Default struct {
					Href   string `json:"href"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Full struct {
					Href string `json:"href"`
				} `json:"full"`
				Wide struct {
					Href string `json:"href"`
				} `json:"wide"`
				Square struct {
					Href string `json:"href"`
				} `json:"square"`
			} `json:"posterImages"`
			Images []struct {
				Name    string `json:"name"`
				URL     string `json:"url"`
				Alt     string `json:"alt"`
				Caption string `json:"caption"`
				Credit  string `json:"credit"`
				Width   int    `json:"width"`
				Height  int    `json:"height"`
			} `json:"images"`
			Thumbnail string `json:"thumbnail"`
			Links     struct {
				API struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Artwork struct {
						Href string `json:"href"`
					} `json:"artwork"`
				} `json:"api"`
				Web struct {
					Href  string `json:"href"`
					Short struct {
						Href string `json:"href"`
					} `json:"short"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"web"`
				Source struct {
					Mezzanine struct {
						Href string `json:"href"`
					} `json:"mezzanine"`
					Flash struct {
						Href string `json:"href"`
					} `json:"flash"`
					Hds struct {
						Href string `json:"href"`
					} `json:"hds"`
					Hls struct {
						Href string `json:"href"`
						Hd   struct {
							Href string `json:"href"`
						} `json:"HD"`
					} `json:"HLS"`
					Hd struct {
						Href string `json:"href"`
					} `json:"HD"`
					Full struct {
						Href string `json:"href"`
					} `json:"full"`
					Href string `json:"href"`
				} `json:"source"`
				Mobile struct {
					Alert struct {
						Href string `json:"href"`
					} `json:"alert"`
					Source struct {
						Href string `json:"href"`
					} `json:"source"`
					Href      string `json:"href"`
					Streaming struct {
						Href string `json:"href"`
					} `json:"streaming"`
					ProgressiveDownload struct {
						Href string `json:"href"`
					} `json:"progressiveDownload"`
				} `json:"mobile"`
			} `json:"links"`
			Title string `json:"title"`
		} `json:"video"`
		Type        string        `json:"type"`
		NowID       string        `json:"nowId"`
		Premium     bool          `json:"premium"`
		Related     []interface{} `json:"related"`
		AllowSearch bool          `json:"allowSearch"`
		Links       struct {
			API struct {
				News struct {
					Href string `json:"href"`
				} `json:"news"`
				Events struct {
					Href string `json:"href"`
				} `json:"events"`
			} `json:"api"`
			Web struct {
				Href  string `json:"href"`
				Short struct {
				} `json:"short"`
			} `json:"web"`
			App struct {
				Sportscenter struct {
					Href string `json:"href"`
				} `json:"sportscenter"`
			} `json:"app"`
			Mobile struct {
				Href string `json:"href"`
			} `json:"mobile"`
		} `json:"links"`
		ID         int `json:"id"`
		Categories []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
			Type        string `json:"type"`
			SportID     int    `json:"sportId"`
			LeagueID    int    `json:"leagueId,omitempty"`
			League      struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Links       struct {
					API struct {
						Leagues struct {
							Href string `json:"href"`
						} `json:"leagues"`
					} `json:"api"`
					Web struct {
						Leagues struct {
							Href string `json:"href"`
						} `json:"leagues"`
					} `json:"web"`
					Mobile struct {
						Leagues struct {
							Href string `json:"href"`
						} `json:"leagues"`
					} `json:"mobile"`
				} `json:"links"`
			} `json:"league,omitempty"`
			UID    string `json:"uid"`
			TeamID int    `json:"teamId,omitempty"`
			Team   struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Links       struct {
					API struct {
						Teams struct {
							Href string `json:"href"`
						} `json:"teams"`
					} `json:"api"`
					Web struct {
						Teams struct {
							Href string `json:"href"`
						} `json:"teams"`
					} `json:"web"`
					Mobile struct {
						Teams struct {
							Href string `json:"href"`
						} `json:"teams"`
					} `json:"mobile"`
				} `json:"links"`
			} `json:"team,omitempty"`
		} `json:"categories"`
		Headline string `json:"headline"`
		GameID   string `json:"gameId"`
		Images   []struct {
			Name    string `json:"name"`
			Width   int    `json:"width"`
			Alt     string `json:"alt"`
			Caption string `json:"caption"`
			URL     string `json:"url"`
			Height  int    `json:"height"`
		} `json:"images"`
		LinkText     string    `json:"linkText"`
		Published    time.Time `json:"published"`
		LastModified time.Time `json:"lastModified"`
		Metrics      []struct {
			Count int    `json:"count"`
			Type  string `json:"type"`
		} `json:"metrics"`
		Inlines []interface{} `json:"inlines"`
		Story   string        `json:"story"`
	} `json:"article"`
	TicketsInfo struct {
		Tickets []struct {
			TicketName string `json:"ticketName"`
			TicketLink string `json:"ticketLink"`
			Type       string `json:"type"`
		} `json:"tickets"`
		SeatSituation struct {
			OpponentTeamName string    `json:"opponentTeamName"`
			CurrentTeamName  string    `json:"currentTeamName"`
			VenueName        string    `json:"venueName"`
			Summary          string    `json:"summary"`
			Date             time.Time `json:"date"`
			DateShort        string    `json:"dateShort"`
			DateDay          string    `json:"dateDay"`
			HomeAway         string    `json:"homeAway"`
			EventLink        string    `json:"eventLink"`
			VenueLink        string    `json:"venueLink"`
			GenericLink      string    `json:"genericLink"`
			TeamLink         string    `json:"teamLink"`
		} `json:"seatSituation"`
	} `json:"ticketsInfo"`
	Standings struct {
		FullViewLink struct {
			Text string `json:"text"`
			Href string `json:"href"`
		} `json:"fullViewLink"`
		Groups []struct {
			Standings struct {
				Entries []struct {
					Team  string `json:"team"`
					Link  string `json:"link"`
					ID    string `json:"id"`
					UID   string `json:"uid"`
					Stats []struct {
						Name             string  `json:"name"`
						DisplayName      string  `json:"displayName,omitempty"`
						ShortDisplayName string  `json:"shortDisplayName,omitempty"`
						Description      string  `json:"description,omitempty"`
						Abbreviation     string  `json:"abbreviation"`
						Type             string  `json:"type"`
						Value            float64 `json:"value,omitempty"`
						DisplayValue     string  `json:"displayValue"`
						ID               string  `json:"id,omitempty"`
						Summary          string  `json:"summary,omitempty"`
					} `json:"stats"`
					Logo []struct {
						Href        string   `json:"href"`
						Width       int      `json:"width"`
						Height      int      `json:"height"`
						Alt         string   `json:"alt"`
						Rel         []string `json:"rel"`
						LastUpdated string   `json:"lastUpdated"`
					} `json:"logo"`
				} `json:"entries"`
			} `json:"standings"`
			Header string `json:"header"`
			Href   string `json:"href"`
		} `json:"groups"`
	} `json:"standings"`
}

type Game struct {
	ID        int
	AwayScore int
	AwayTeam  string
	HomeScore int
	HomeTeam  string
	Quarter   int
	Time      string
	GameOver  bool
	Odds      struct {
		Details   string
		OverUnder float64
	}
	Digits struct {
		Away1 int
		Away2 int
		Away3 int
		Away4 int
		Home1 int
		Home2 int
		Home3 int
		Home4 int
	}
}
