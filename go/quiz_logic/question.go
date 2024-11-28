package quiz_logic

import (
	"fmt"
	"strconv"
	"strings"
)

// Question interface defines the common behavior for all question types
type Question interface {
	getQuestion() string
	getType() string
	checkAnswer(answer string) bool
	getOptions() []string
}

// BaseQuestion contains common fields for all question types
type BaseQuestion struct {
	QuestionText string   `json:"question"`
	Type         string   `json:"type"`
	Answers      []string `json:"answers"`
}

func (bq *BaseQuestion) getQuestion() string {
	return bq.QuestionText
}

func (bq *BaseQuestion) getType() string {
	return bq.Type
}

// MultipleChoiceQuestion implements Question interface
type MultipleChoiceQuestion struct {
	BaseQuestion
	Options []string `json:"options"`
}

func (mcq *MultipleChoiceQuestion) checkAnswer(answer string) bool {
	// First try to parse the answer as a number
	if num, err := strconv.Atoi(answer); err == nil && num > 0 && num <= len(mcq.Options) {
		// Convert to zero-based index
		answer = mcq.Options[num-1]
	}

	// Convert answer to lowercase for case-insensitive comparison
	answer = strings.ToLower(answer)

	// Check if the answer matches any of the correct answers
	for _, correctAnswer := range mcq.Answers {
		if answer == strings.ToLower(correctAnswer) {
			return true
		}
	}
	return false
}

func (mcq *MultipleChoiceQuestion) getOptions() []string {
	return mcq.Options
}

// TrueFalseQuestion implements Question interface
type TrueFalseQuestion struct {
	BaseQuestion
}

func (tfq *TrueFalseQuestion) checkAnswer(answer string) bool {
	// Convert numeric answers to text
	switch answer {
	case "1":
		answer = "true"
	case "2":
		answer = "false"
	}

	// Convert answer to lowercase for case-insensitive comparison
	answer = strings.ToLower(answer)

	// Check if the answer matches any of the correct answers
	for _, correctAnswer := range tfq.Answers {
		if answer == strings.ToLower(correctAnswer) {
			return true
		}
	}
	return false
}

func (tfq *TrueFalseQuestion) getOptions() []string {
	return []string{"True", "False"}
}

// FillInBlankQuestion implements Question interface
type FillInBlankQuestion struct {
	BaseQuestion
}

func (fib *FillInBlankQuestion) checkAnswer(answer string) bool {
	// Convert answer to lowercase for case-insensitive comparison
	answer = strings.ToLower(answer)

	for _, correctAnswer := range fib.Answers {
		if answer == strings.ToLower(correctAnswer) {
			return true
		}
	}
	return false
}

func (fib *FillInBlankQuestion) getOptions() []string {
	return nil
}

// createQuestion factory function to create the appropriate question type
func createQuestion(questionData map[string]interface{}) (Question, error) {
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
