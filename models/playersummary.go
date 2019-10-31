package models

type PlayerSummary struct {
	Fixtures       []PlayerFixture        `json:"fixtures"`
	FixtureHistory []PlayerFixtureHistory `json:"history"`
	History        []PlayerHistory        `json:"history_past"`
}

// fixtures[]
type PlayerFixture struct {
	Code                 int    `json:"code"`
	TeamHome             int    `json:"team_h"`
	TeamHomeScore        int    `json:"team_h_score"`
	TeamAway             int    `json:"team_a"`
	TeamAwayScore        int    `json:"team_a_score"`
	Event                int    `json:"event"`
	Finished             bool   `json:"finished"`
	Minutes              int    `json:"minutes"`
	ProvisionalStartTime bool   `json:"provisional_start_time"`
	KickoffTime          string `json:"kickoff_time"`
	EventName            string `json:"event_name"`
	IsHome               bool   `json:"is_home"`
	Difficulty           int    `json:"difficulty"`
}

// history[]
type PlayerFixtureHistory struct {
	Element          int    `json:"element"`
	Fixture          int    `json:"fixture"`
	OpponentTeam     int    `json:"opponent_team"`
	TotalPoints      int    `json:"total_points"`
	WasHome          bool   `json:"was_home"`
	KickoffTime      string `json:"kickoff_time"`
	TeamHomeScore    int    `json:"team_h_score"`
	TeamAwayScore    int    `json:"team_a_score"`
	Round            int    `json:"round"`
	Minutes          int    `json:"minutes"`
	GoalsScored      int    `json:"goals_scored"`
	Assists          int    `json:"assists"`
	CleanSheets      int    `json:"clean_sheets"`
	GoalsConceded    int    `json:"goals_conceded"`
	OwnGoals         int    `json:"own_goals"`
	PenaltiesSaved   int    `json:"penalties_saved"`
	PenaltiesMissed  int    `json:"penalties_missed"`
	YellowCards      int    `json:"yellow_cards"`
	RedCards         int    `json:"red_cards"`
	Saves            int    `json:"saves"`
	Bonus            int    `json:"bonus"`
	BPS              int    `json:"bps"`
	Influence        string `json:"influence"`
	Creativity       string `json:"creativity"`
	Threat           string `json:"threat"`
	ICTIndex         string `json:"ict_index"`
	Value            int    `json:"value"`
	TransfersBalance int    `json:"transfers_balance"`
	Selected         int    `json:"selected"`
	TransfersIn      int    `json:"transfers_in"`
	TransfersOut     int    `json:"transfers_out"`
}

// history_past[]
type PlayerHistory struct {
	SeasonName      string `json:"season_name"`
	ElementCode     int    `json:"element_code"`
	StartCost       int    `json:"start_cost"`
	EndCost         int    `json:"end_cost"`
	TotalPoints     int    `json:"total_points"`
	Minutes         int    `json:"minutes"`
	GoalsScored     int    `json:"goals_scored"`
	Assists         int    `json:"assists"`
	CleanSheets     int    `json:"clean_sheets"`
	GoalsConceded   int    `json:"goals_conceded"`
	OwnGoals        int    `json:"own_goals"`
	PenaltiesSaved  int    `json:"penalties_saved"`
	PenaltiesMissed int    `json:"penalties_missed"`
	YellowCards     int    `json:"yellow_cards"`
	RedCards        int    `json:"red_cards"`
	Saves           int    `json:"saves"`
	Bonus           int    `json:"bonus"`
	BPS             int    `json:"bps"`
	Influence       string `json:"influence"`
	Creativity      string `json:"creativity"`
	Threat          string `json:"thread"`
	ICTIndex        string `json:"ict_index"`
}
