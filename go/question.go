package main

import "fmt"

// Question interface defines the common behavior for all question types
type Question interface {
	GetQuestion() string
	GetType() string
	CheckAnswer(answer string) bool
	GetOptions() []string
}

// BaseQuestion contains common fields for all question types
type BaseQuestion struct {
	QuestionText string   `json:"question"`
	Type         string   `json:"type"`
	Answers      []string `json:"answers"`
}

func (bq *BaseQuestion) GetQuestion() string {
	return bq.QuestionText
}

func (bq *BaseQuestion) GetType() string {
	return bq.Type
}

// MultipleChoiceQuestion implements Question interface
type MultipleChoiceQuestion struct {
	BaseQuestion
	Options []string `json:"options"`
}

func (mcq *MultipleChoiceQuestion) CheckAnswer(answer string) bool {
	for _, correctAnswer := range mcq.Answers {
		if answer == correctAnswer {
			return true
		}
	}
	return false
}

func (mcq *MultipleChoiceQuestion) GetOptions() []string {
	return mcq.Options
}

// TrueFalseQuestion implements Question interface
type TrueFalseQuestion struct {
	BaseQuestion
}

func (tfq *TrueFalseQuestion) CheckAnswer(answer string) bool {
	for _, correctAnswer := range tfq.Answers {
		if answer == correctAnswer {
			return true
		}
	}
	return false
}

func (tfq *TrueFalseQuestion) GetOptions() []string {
	return []string{"True", "False"}
}

// FillInBlankQuestion implements Question interface
type FillInBlankQuestion struct {
	BaseQuestion
}

func (fib *FillInBlankQuestion) CheckAnswer(answer string) bool {
	for _, correctAnswer := range fib.Answers {
		if answer == correctAnswer {
			return true
		}
	}
	return false
}

func (fib *FillInBlankQuestion) GetOptions() []string {
	return nil
}

// CreateQuestion factory function to create the appropriate question type
func CreateQuestion(questionData map[string]interface{}) (Question, error) {
	// Extract common fields
	baseQuestion := BaseQuestion{
		QuestionText: questionData["question"].(string),
		Type:         questionData["type"].(string),
	}

	// Extract answers
	if answersData, ok := questionData["answers"].([]interface{}); ok {
		answers := make([]string, len(answersData))
		for i, ans := range answersData {
			answers[i] = ans.(string)
		}
		baseQuestion.Answers = answers
	}

	// Create specific question type
	switch baseQuestion.Type {
	case "multiple_choice":
		mcq := &MultipleChoiceQuestion{BaseQuestion: baseQuestion}
		if optionsData, ok := questionData["options"].([]interface{}); ok {
			options := make([]string, len(optionsData))
			for i, opt := range optionsData {
				options[i] = opt.(string)
			}
			mcq.Options = options
		}
		return mcq, nil
	case "true_false":
		return &TrueFalseQuestion{BaseQuestion: baseQuestion}, nil
	case "fill_in_blank":
		return &FillInBlankQuestion{BaseQuestion: baseQuestion}, nil
	default:
		return nil, fmt.Errorf("unknown question type: %s", baseQuestion.Type)
	}
}
