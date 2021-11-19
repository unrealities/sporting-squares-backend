package nflstats

// ESPN NFL Scores
type EspnNFLScores struct {
	Events []struct {
		Competitions []struct {
			Attendance int `json:"attendance"`
			Broadcasts []struct {
				Market string   `json:"market"`
				Names  []string `json:"names"`
			} `json:"broadcasts"`
			Competitors []struct {
				HomeAway   string `json:"homeAway"`
				ID         string `json:"id"`
				Linescores []struct {
					Value int `json:"value"`
				} `json:"linescores"`
				Order   int `json:"order"`
				Records []struct {
					Abbreviation string `json:"abbreviation,omitempty"`
					Name         string `json:"name"`
					Summary      string `json:"summary"`
					Type         string `json:"type"`
				} `json:"records"`
				Score      string        `json:"score"`
				Statistics []interface{} `json:"statistics"`
				Team       struct {
					Abbreviation   string `json:"abbreviation"`
					AlternateColor string `json:"alternateColor"`
					Color          string `json:"color"`
					DisplayName    string `json:"displayName"`
					ID             string `json:"id"`
					IsActive       bool   `json:"isActive"`
					Links          []struct {
						Href       string   `json:"href"`
						IsExternal bool     `json:"isExternal"`
						IsPremium  bool     `json:"isPremium"`
						Rel        []string `json:"rel"`
						Text       string   `json:"text"`
					} `json:"links"`
					Location         string `json:"location"`
					Logo             string `json:"logo"`
					Name             string `json:"name"`
					ShortDisplayName string `json:"shortDisplayName"`
					UID              string `json:"uid"`
					Venue            struct {
						ID string `json:"id"`
					} `json:"venue"`
				} `json:"team"`
				Type   string `json:"type"`
				UID    string `json:"uid"`
				Winner bool   `json:"winner"`
			} `json:"competitors"`
			ConferenceCompetition bool   `json:"conferenceCompetition"`
			Date                  string `json:"date"`
			GeoBroadcasts         []struct {
				Lang   string `json:"lang"`
				Market struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"market"`
				Media struct {
					ShortName string `json:"shortName"`
				} `json:"media"`
				Region string `json:"region"`
				Type   struct {
					ID        string `json:"id"`
					ShortName string `json:"shortName"`
				} `json:"type"`
			} `json:"geoBroadcasts"`
			Headlines []struct {
				Description   string `json:"description"`
				ShortLinkText string `json:"shortLinkText"`
				Type          string `json:"type"`
			} `json:"headlines"`
			ID      string `json:"id"`
			Leaders []struct {
				Abbreviation string `json:"abbreviation"`
				DisplayName  string `json:"displayName"`
				Leaders      []struct {
					Athlete struct {
						Active      bool   `json:"active"`
						DisplayName string `json:"displayName"`
						FullName    string `json:"fullName"`
						Headshot    string `json:"headshot"`
						ID          string `json:"id"`
						Jersey      string `json:"jersey"`
						Links       []struct {
							Href string   `json:"href"`
							Rel  []string `json:"rel"`
						} `json:"links"`
						Position struct {
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
						ShortName string `json:"shortName"`
						Team      struct {
							ID string `json:"id"`
						} `json:"team"`
					} `json:"athlete"`
					DisplayValue string `json:"displayValue"`
					Team         struct {
						ID string `json:"id"`
					} `json:"team"`
					Value int `json:"value"`
				} `json:"leaders"`
				Name             string `json:"name"`
				ShortDisplayName string `json:"shortDisplayName"`
			} `json:"leaders"`
			NeutralSite bool          `json:"neutralSite"`
			Notes       []interface{} `json:"notes"`
			Recent      bool          `json:"recent"`
			StartDate   string        `json:"startDate"`
			Status      struct {
				Clock        int    `json:"clock"`
				DisplayClock string `json:"displayClock"`
				Period       int    `json:"period"`
				Type         struct {
					Completed   bool   `json:"completed"`
					Description string `json:"description"`
					Detail      string `json:"detail"`
					ID          string `json:"id"`
					Name        string `json:"name"`
					ShortDetail string `json:"shortDetail"`
					State       string `json:"state"`
				} `json:"type"`
			} `json:"status"`
			TimeValid bool `json:"timeValid"`
			Type      struct {
				Abbreviation string `json:"abbreviation"`
				ID           string `json:"id"`
			} `json:"type"`
			UID   string `json:"uid"`
			Venue struct {
				Address struct {
					City  string `json:"city"`
					State string `json:"state"`
				} `json:"address"`
				Capacity int    `json:"capacity"`
				FullName string `json:"fullName"`
				ID       string `json:"id"`
				Indoor   bool   `json:"indoor"`
			} `json:"venue"`
		} `json:"competitions"`
		Date  string `json:"date"`
		ID    string `json:"id"`
		Links []struct {
			Href       string   `json:"href"`
			IsExternal bool     `json:"isExternal"`
			IsPremium  bool     `json:"isPremium"`
			Language   string   `json:"language"`
			Rel        []string `json:"rel"`
			ShortText  string   `json:"shortText"`
			Text       string   `json:"text"`
		} `json:"links"`
		Name   string `json:"name"`
		Season struct {
			Slug string `json:"slug"`
			Type int    `json:"type"`
			Year int    `json:"year"`
		} `json:"season"`
		ShortName string `json:"shortName"`
		Status    struct {
			Clock        int    `json:"clock"`
			DisplayClock string `json:"displayClock"`
			Period       int    `json:"period"`
			Type         struct {
				Completed   bool   `json:"completed"`
				Description string `json:"description"`
				Detail      string `json:"detail"`
				ID          string `json:"id"`
				Name        string `json:"name"`
				ShortDetail string `json:"shortDetail"`
				State       string `json:"state"`
			} `json:"type"`
		} `json:"status"`
		UID string `json:"uid"`
	} `json:"events"`
	Leagues []struct {
		Abbreviation string `json:"abbreviation"`
		Calendar     []struct {
			EndDate string `json:"endDate"`
			Entries []struct {
				AlternateLabel string `json:"alternateLabel"`
				Detail         string `json:"detail"`
				EndDate        string `json:"endDate"`
				Label          string `json:"label"`
				StartDate      string `json:"startDate"`
				Value          string `json:"value"`
			} `json:"entries"`
			Label     string `json:"label"`
			StartDate string `json:"startDate"`
			Value     string `json:"value"`
		} `json:"calendar"`
		CalendarEndDate     string `json:"calendarEndDate"`
		CalendarIsWhitelist bool   `json:"calendarIsWhitelist"`
		CalendarStartDate   string `json:"calendarStartDate"`
		CalendarType        string `json:"calendarType"`
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Season              struct {
			EndDate   string `json:"endDate"`
			StartDate string `json:"startDate"`
			Type      struct {
				Abbreviation string `json:"abbreviation"`
				ID           string `json:"id"`
				Name         string `json:"name"`
				Type         int    `json:"type"`
			} `json:"type"`
			Year int `json:"year"`
		} `json:"season"`
		Slug string `json:"slug"`
		UID  string `json:"uid"`
	} `json:"leagues"`
	Season struct {
		Type int `json:"type"`
		Year int `json:"year"`
	} `json:"season"`
	Week struct {
		Number     int `json:"number"`
		TeamsOnBye []struct {
			Abbreviation string `json:"abbreviation"`
			DisplayName  string `json:"displayName"`
			ID           string `json:"id"`
			IsActive     bool   `json:"isActive"`
			Links        []struct {
				Href       string   `json:"href"`
				IsExternal bool     `json:"isExternal"`
				IsPremium  bool     `json:"isPremium"`
				Rel        []string `json:"rel"`
				Text       string   `json:"text"`
			} `json:"links"`
			Location         string `json:"location"`
			Logo             string `json:"logo"`
			Name             string `json:"name"`
			ShortDisplayName string `json:"shortDisplayName"`
			UID              string `json:"uid"`
		} `json:"teamsOnBye"`
	} `json:"week"`
}

type Scores struct{}
