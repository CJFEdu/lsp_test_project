package quiz_logic

import (
	"testing"
)

func TestMultipleChoiceQuestion(t *testing.T) {
	mcq := &MultipleChoiceQuestion{
		BaseQuestion: BaseQuestion{
			QuestionText: "What is the capital of France?",
			Type:         "multiple_choice",
			Answers:      []string{"Paris"},
		},
		Options: []string{"London", "Paris", "Berlin", "Madrid"},
	}

	tests := []struct {
		name     string
		answer   string
		expected bool
	}{
		{"Correct text answer", "Paris", true},
		{"Correct numeric answer", "2", true},
		{"Wrong text answer", "London", false},
		{"Wrong numeric answer", "1", false},
		{"Invalid numeric answer", "5", false},
		{"Empty answer", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mcq.checkAnswer(tt.answer); got != tt.expected {
				t.Errorf("MultipleChoiceQuestion.CheckAnswer(%q) = %v, want %v", tt.answer, got, tt.expected)
			}
		})
	}

	// Test GetQuestion
	if got := mcq.getQuestion(); got != "What is the capital of France?" {
		t.Errorf("GetQuestion() = %v, want %v", got, "What is the capital of France?")
	}

	// Test GetType
	if got := mcq.getType(); got != "multiple_choice" {
		t.Errorf("GetType() = %v, want %v", got, "multiple_choice")
	}

	// Test GetOptions
	options := mcq.getOptions()
	if len(options) != 4 {
		t.Errorf("GetOptions() returned %d options, want 4", len(options))
	}
}

func TestTrueFalseQuestion(t *testing.T) {
	tfq := &TrueFalseQuestion{
		BaseQuestion: BaseQuestion{
			QuestionText: "Is Paris the capital of France?",
			Type:         "true_false",
			Answers:      []string{"True"},
		},
	}

	tests := []struct {
		name     string
		answer   string
		expected bool
	}{
		{"Correct text answer", "True", true},
		{"Correct numeric answer", "1", true},
		{"Wrong text answer", "False", false},
		{"Wrong numeric answer", "2", false},
		{"Invalid answer", "Maybe", false},
		{"Empty answer", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tfq.checkAnswer(tt.answer); got != tt.expected {
				t.Errorf("TrueFalseQuestion.CheckAnswer(%q) = %v, want %v", tt.answer, got, tt.expected)
			}
		})
	}

	options := tfq.getOptions()
	if len(options) != 2 || options[0] != "True" || options[1] != "False" {
		t.Errorf("GetOptions() = %v, want [True False]", options)
	}
}

func TestFillInBlankQuestion(t *testing.T) {
	fib := &FillInBlankQuestion{
		BaseQuestion: BaseQuestion{
			QuestionText: "The capital of France is ___.",
			Type:         "fill_in_blank",
			Answers:      []string{"Paris"},
		},
	}

	tests := []struct {
		name     string
		answer   string
		expected bool
	}{
		{"Correct answer", "Paris", true},
		{"Wrong answer", "London", false},
		{"Empty answer", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib.checkAnswer(tt.answer); got != tt.expected {
				t.Errorf("FillInBlankQuestion.checkAnswer(%q) = %v, want %v", tt.answer, got, tt.expected)
			}
		})
	}

	if options := fib.getOptions(); options != nil {
		t.Errorf("GetOptions() = %v, want nil", options)
	}
}

func TestCreateQuestion(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Create multiple choice question",
			data: map[string]interface{}{
				"question": "What is the capital of France?",
				"type":     "multiple_choice",
				"answers":  []interface{}{"Paris"},
				"options":  []interface{}{"London", "Paris", "Berlin", "Madrid"},
			},
			wantErr: false,
		},
		{
			name: "Create true/false question",
			data: map[string]interface{}{
				"question": "Is Paris the capital of France?",
				"type":     "true_false",
				"answers":  []interface{}{"True"},
			},
			wantErr: false,
		},
		{
			name: "Create fill in blank question",
			data: map[string]interface{}{
				"question": "The capital of France is ___.",
				"type":     "fill_in_blank",
				"answers":  []interface{}{"Paris"},
			},
			wantErr: false,
		},
		{
			name: "Invalid question type",
			data: map[string]interface{}{
				"question": "Test question",
				"type":     "invalid_type",
				"answers":  []interface{}{"Test"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, err := createQuestion(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && q == nil {
				t.Error("CreateQuestion() returned nil question without error")
			}
		})
	}
}
