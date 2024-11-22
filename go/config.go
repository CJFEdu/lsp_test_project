package main

// Config represents the quiz configuration
type Config struct {
	Title          string     `json:"title"`
	TimeLimit      int        `json:"timeLimit"`
	RandomizeOrder bool       `json:"randomizeOrder"`
	PassingScore   int        `json:"passingScore"`
	Questions      [][]string `json:"questions"`
	Settings       struct {
		ShowFeedbackAfterEach bool `json:"showFeedbackAfterEach"`
		AllowSkipping         bool `json:"allowSkipping"`
		ShowTimer             bool `json:"showTimer"`
	} `json:"settings"`
}
