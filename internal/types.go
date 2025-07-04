package internal

/* DB data shape */

/* ESPN NBA API Payloads */

type teamFetchPayload struct {
	Sports []struct {
		Leagues []struct {
			Teams []struct {
				Team Team `json:"team"`
			} `json:"teams"`
		} `json:"leagues"`
	} `json:"sports"`
}

type rosterFetchPayload struct {
	Athletes []Player `json:"athletes"`
}

type seasonInfoPayload struct {
	Leagues []struct {
		SeasonInfo struct {
			Year      int    `json:"year"`
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		} `json:"season"`
	} `json:"leagues"`
}

type seasonGamesFetchPayload struct {
	Events []rawGameWrapper `json:"events"`
}

// Needed bc playByPlay is nested in event, but we want it flat in Game obj
type rawGameWrapper struct {
	Game
	Competitions []struct {
		PlayByPlayAvailable bool `json:"playByPlayAvailable"`
	} `json:"competitions"`
}

type gameDataFetchPayload struct {
	Plays []play `json:"plays"`
}

/* Team & Player  */

type Team struct {
	Id             string `json:"id"`
	Uid            string `json:"uid"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Color          string `json:"color"`
	AlternateColor string `json:"alternateColor"`
	Logos          []struct {
		Href string `json:"href"`
	} `json:"logos"`
	Roster []Player `json:"roster"`
}

type Player struct {
	Id          string       `json:"id"`
	Uid         string       `json:"uid"`
	Guid        string       `json:"guid"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	FullName    string       `json:"fullName"`
	DisplayName string       `json:"displayName"`
	ShortName   string       `json:"shortName"`
	Slug        string       `json:"slug"`
	Status      playerStatus `json:"status"`
	BetStats    betStats     `json:"betStats"`
	Headshot    struct {
		Href string `json:"href"`
	} `json:"headshot"`
}

type playerStatus struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type betStats struct {
	TipOffWinPer        uint8 `json:"tipOffWinPer"`
	FirstPointsMade     uint8 `json:"firstPointsMade"`
	FirstShotAttempts   uint8 `json:"firstShotAttempts"`
	ScoreOnFirstAttempt uint8 `json:"scoreOnFirstAttempt"`
}

/* Game */

type Game struct {
	Id            string        `json:"id"`
	Uid           string        `json:"uid"`
	Date          string        `json:"date"`
	Name          string        `json:"name"`
	Season        Season        `json:"season"`
	PlayByPlay    bool          `json:"playByPlay"`
	TrackedEvents trackedEvents `json:"trackedEvents"`
	Players       []gamePlayer  `json:"gamePlayers"`
}

type Season struct {
	Year int    `json:"year"`
	Type int    `json:"type"`
	Slug string `json:"slug"`
}

/* Game Player & Athlete */

type gamePlayer struct {
	TeamId  int     `json:"teamId"`
	Active  bool    `json:"active"`
	Athlete athlete `json:"athlete"`
}

type athlete struct {
	Id          string `json:"id"`
	Guid        string `json:"guid"`
	DisplayName string `json:"displayName"`
	DidNotPlay  bool   `json:"didNotPlay"`
}

/* Play */

type trackedEvents struct {
	TipOff           play `json:"tipOff"`
	FirstScore       play `json:"firstScore"`
	FirstShotAttempt play `json:"firstShotAttempt"`
}

type play struct {
	Id string `json:"id"`

	Type struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"type"`

	Text      string `json:"text"`
	AwayScore uint8  `json:"awayScore"`
	HomeScore uint8  `json:"homeScore"`

	Period struct {
		Number       uint8  `json:"number"`
		DisplayValue string `json:"displayValue"`
	} `json:"period"`

	Clock struct {
		DisplayValue string `json:"displayValue"`
	} `json:"clock"`

	ScoringPlay  bool  `json:"scoringPlay"`
	ScoreValue   uint8 `json:"scoreValue"`
	ShootingPlay bool  `json:"shootingPlay"`

	Team struct {
		Id string `json:"id"`
	} `json:"team"`

	Participants []Participant `json:"participants"`
}

type Participant struct {
	Athlete pAthlete `json:"athlete"`
}

type pAthlete struct {
	Id string `json:"id"`
}
