package main

// Config represents the quiz configuration
type Config struct {
	Title        string `json:"title"`
	TimeLimit    int    `json:"timelimit"`    // in minutes
	Randomize    bool   `json:"randomize"`    // randomize question order
	PassingScore int    `json:"passingScore"` // percentage needed to pass
}
