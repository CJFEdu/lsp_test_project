package quiz_logic

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Quiz struct {
	Config         Config
	Questions      []Question
	startTime      time.Time
	correctAnswers int
	totalQuestions int
}

func (q *Quiz) selectQuestions(quizPath string) error {
	// Create a map to track loaded questions
	loadedQuestions := make(map[string]Question)

	// First, load all question files and store them in the map
	files, err := os.ReadDir(quizPath)
	if err != nil {
		return fmt.Errorf("error reading quiz directory: %v", err)
	}

	for _, file := range files {
		if file.Name() == "config.json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(quizPath, file.Name()))
		if err != nil {
			return fmt.Errorf("error reading question file %s: %v", file.Name(), err)
		}

		var questionData map[string]interface{}
		if err := json.Unmarshal(data, &questionData); err != nil {
			return fmt.Errorf("error parsing question file %s: %v", file.Name(), err)
		}

		question, err := createQuestion(questionData)
		if err != nil {
			return fmt.Errorf("error creating question from %s: %v", file.Name(), err)
		}

		// Store the question in the map using filename without extension as key
		key := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		loadedQuestions[key] = question
	}

	// Now process the question sets from the config
	for _, questionSet := range q.Config.Questions {
		if len(questionSet) == 0 {
			continue // Skip empty sets
		}

		// Randomly select one question from the set
		questionID := questionSet[rand.Intn(len(questionSet))]

		if question, exists := loadedQuestions[questionID]; exists {
			q.Questions = append(q.Questions, question)
		} else {
			return fmt.Errorf("question file not found: %s", questionID)
		}
	}

	if q.Config.RandomizeOrder {
		rand.Shuffle(len(q.Questions), func(i, j int) {
			q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
		})
	}

	return nil
}

func (q *Quiz) Run() {
	q.startTime = time.Now()
	q.totalQuestions = len(q.Questions)
	q.correctAnswers = 0

	fmt.Printf("\nStarting Quiz: %s\n", q.Config.Title)
	if q.Config.TimeLimit > 0 && q.Config.Settings.ShowTimer {
		fmt.Printf("Time Limit: %d minutes\n", q.Config.TimeLimit)
	}
	fmt.Printf("Number of Questions: %d\n\n", len(q.Questions))

	for i, question := range q.Questions {
		if q.Config.Settings.ShowTimer && q.isTimeUp() {
			fmt.Println("\nTime's up!")
			break
		}

		fmt.Printf("\nQuestion %d: %s\n", i+1, question.getQuestion())
		options := question.getOptions()
		if len(options) > 0 {
			fmt.Println("Options:")
			for j, option := range options {
				fmt.Printf("%d. %s\n", j+1, option)
			}
		}

		var answer string
		if q.Config.Settings.AllowSkipping {
			fmt.Print("\nEnter your answer (or press Enter to skip): ")
		} else {
			fmt.Print("\nEnter your answer: ")
		}
		fmt.Scanln(&answer)

		if answer == "" && q.Config.Settings.AllowSkipping {
			fmt.Println("Question skipped.")
			continue
		}

		if question.checkAnswer(answer) {
			q.correctAnswers++
			if q.Config.Settings.ShowFeedbackAfterEach {
				fmt.Println("Correct!")
			}
		} else if q.Config.Settings.ShowFeedbackAfterEach {
			fmt.Println("Incorrect.")
		}
	}

	score := q.calculateScore()
	fmt.Printf("\nQuiz completed!\nScore: %d/%d (%d%%)\n", q.correctAnswers, q.totalQuestions, score)
	if q.hasPassed() {
		fmt.Println("Congratulations! You passed!")
	} else {
		fmt.Println("Sorry, you didn't pass. Keep practicing!")
	}
}

func (q *Quiz) calculateScore() int {
	if q.totalQuestions == 0 {
		return 0
	}
	return (q.correctAnswers * 100) / q.totalQuestions
}

func (q *Quiz) isTimeUp() bool {
	if q.Config.TimeLimit == 0 {
		return false
	}
	elapsed := time.Since(q.startTime)
	return int(elapsed.Seconds()) >= q.Config.TimeLimit*60
}

func (q *Quiz) hasPassed() bool {
	if q.Config.PassingScore == 0 {
		return true // If no passing score is set, consider it passed
	}
	return q.calculateScore() >= q.Config.PassingScore
}
