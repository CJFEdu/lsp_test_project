package quiz_logic

// Config represents the quiz configuration
type Config struct {
	Title          string     `json:"title"`
	TimeLimit      int        `json:"timeLimit"`      // in minutes
	RandomizeOrder bool       `json:"randomizeOrder"` // randomize question order
	PassingScore   int        `json:"passingScore"`   // percentage needed to pass
	Questions      [][]string `json:"questions"`      // list of question sets
	Settings       struct {
		ShowFeedbackAfterEach bool `json:"showFeedbackAfterEach"`
		AllowSkipping         bool `json:"allowSkipping"`
		ShowTimer             bool `json:"showTimer"`
	} `json:"settings"`
}
