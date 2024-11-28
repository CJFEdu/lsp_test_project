package quiz_logic

import (
	"testing"
	"time"
)

func TestQuiz_SelectQuestions(t *testing.T) {
	quiz := &Quiz{
		Config: Config{
			Title:     "Test Quiz",
			TimeLimit: 300,
		},
	}

	// Create test questions
	mcq := &MultipleChoiceQuestion{
		BaseQuestion: BaseQuestion{
			QuestionText: "What is the capital of France?",
			Type:         "multiple_choice",
			Answers:      []string{"Paris"},
		},
		Options: []string{"London", "Paris", "Berlin", "Madrid"},
	}

	tfq := &TrueFalseQuestion{
		BaseQuestion: BaseQuestion{
			QuestionText: "Is Paris in France?",
			Type:         "true_false",
			Answers:      []string{"True"},
		},
	}

	// Add questions to quiz
	quiz.Questions = []Question{mcq, tfq}

	// Test question handling
	if len(quiz.Questions) != 2 {
		t.Errorf("Expected 2 questions, got %d", len(quiz.Questions))
	}

	// Verify question types
	if quiz.Questions[0].getType() != "multiple_choice" {
		t.Errorf("Expected multiple_choice question, got %s", quiz.Questions[0].getType())
	}

	if quiz.Questions[1].getType() != "true_false" {
		t.Errorf("Expected true_false question, got %s", quiz.Questions[1].getType())
	}

	// Test answer checking
	if !quiz.Questions[0].checkAnswer("Paris") {
		t.Error("Failed to recognize correct answer 'Paris'")
	}

	if !quiz.Questions[0].checkAnswer("2") {
		t.Error("Failed to recognize correct numeric answer '2'")
	}

	if quiz.Questions[0].checkAnswer("London") {
		t.Error("Incorrectly accepted wrong answer 'London'")
	}
}

func TestQuiz_CalculateScore(t *testing.T) {
	tests := []struct {
		name           string
		correctAnswers int
		totalQuestions int
		expectedScore  int
	}{
		{
			name:           "Perfect score",
			correctAnswers: 10,
			totalQuestions: 10,
			expectedScore:  100,
		},
		{
			name:           "70% score",
			correctAnswers: 7,
			totalQuestions: 10,
			expectedScore:  70,
		},
		{
			name:           "Zero score",
			correctAnswers: 0,
			totalQuestions: 10,
			expectedScore:  0,
		},
		{
			name:           "No questions",
			correctAnswers: 0,
			totalQuestions: 0,
			expectedScore:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quiz := &Quiz{
				correctAnswers: tt.correctAnswers,
				totalQuestions: tt.totalQuestions,
			}

			score := quiz.calculateScore()
			if score != tt.expectedScore {
				t.Errorf("calculateScore() = %v, want %v", score, tt.expectedScore)
			}
		})
	}
}

func TestQuiz_IsTimeUp(t *testing.T) {
	tests := []struct {
		name      string
		timeLimit int
		sleep     time.Duration
		want      bool
	}{
		{
			name:      "Time not up",
			timeLimit: 2,
			sleep:     time.Second,
			want:      false,
		},
		{
			name:      "Time is up",
			timeLimit: 1,
			sleep:     2 * time.Second,
			want:      true,
		},
		{
			name:      "No time limit",
			timeLimit: 0,
			sleep:     time.Second,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quiz := &Quiz{
				Config: Config{
					TimeLimit: tt.timeLimit,
				},
			}
			quiz.Run()
			time.Sleep(tt.sleep)
			if got := quiz.isTimeUp(); got != tt.want {
				t.Errorf("IsTimeUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuiz_HasPassed(t *testing.T) {
	tests := []struct {
		name           string
		passScore      int
		correctAnswers int
		totalQuestions int
		want           bool
	}{
		{
			name:           "Pass with exact score",
			passScore:      70,
			correctAnswers: 7,
			totalQuestions: 10,
			want:           true,
		},
		{
			name:           "Pass with higher score",
			passScore:      70,
			correctAnswers: 8,
			totalQuestions: 10,
			want:           true,
		},
		{
			name:           "Fail",
			passScore:      70,
			correctAnswers: 6,
			totalQuestions: 10,
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quiz := &Quiz{
				Config: Config{
					PassingScore: tt.passScore,
				},
				correctAnswers: tt.correctAnswers,
				totalQuestions: tt.totalQuestions,
			}

			if got := quiz.hasPassed(); got != tt.want {
				t.Errorf("hasPassed() = %v, want %v", got, tt.want)
			}
		})
	}
}
