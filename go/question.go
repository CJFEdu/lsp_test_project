package main

// Question represents a single quiz question
type Question struct {
	Question string   `json:"question"`
	Type     string   `json:"type"`
	Options  []string `json:"options,omitempty"`
	Answers  []string `json:"answers"`
}
