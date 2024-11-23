package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Quiz struct {
	Config     Config
	Questions  []Question
	TimeLimit  time.Duration
	StartTime  time.Time
	Score      int
	TotalScore int
}

func (q *Quiz) selectQuestions(quizPath string) error {
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

		question, err := CreateQuestion(questionData)
		if err != nil {
			return fmt.Errorf("error creating question from %s: %v", file.Name(), err)
		}

		q.Questions = append(q.Questions, question)
	}

	if q.Config.Randomize {
		rand.Shuffle(len(q.Questions), func(i, j int) {
			q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
		})
	}

	return nil
}

func (q *Quiz) run() {
	q.StartTime = time.Now()
	q.TotalScore = len(q.Questions)

	fmt.Printf("\nStarting Quiz: %s\n", q.Config.Title)
	if q.Config.TimeLimit > 0 {
		fmt.Printf("Time Limit: %d minutes\n", q.Config.TimeLimit)
	}
	fmt.Printf("Number of Questions: %d\n\n", len(q.Questions))

	for i, question := range q.Questions {
		if q.Config.TimeLimit > 0 {
			elapsed := time.Since(q.StartTime)
			if elapsed >= time.Duration(q.Config.TimeLimit)*time.Minute {
				fmt.Println("\nTime's up!")
				break
			}
			fmt.Printf("Time remaining: %d minutes\n", q.Config.TimeLimit-int(elapsed.Minutes()))
		}

		fmt.Printf("\nQuestion %d: %s\n", i+1, question.GetQuestion())
		
		options := question.GetOptions()
		if options != nil {
			fmt.Println("Options:")
			for j, option := range options {
				fmt.Printf("%d. %s\n", j+1, option)
			}
		}

		var answer string
		fmt.Print("Your answer: ")
		fmt.Scanln(&answer)

		if question.CheckAnswer(answer) {
			fmt.Println("Correct!")
			q.Score++
		} else {
			fmt.Println("Incorrect.")
		}
	}

	fmt.Printf("\nQuiz Complete!\n")
	fmt.Printf("Score: %d/%d (%.1f%%)\n", q.Score, q.TotalScore, float64(q.Score)/float64(q.TotalScore)*100)

	if q.Config.PassingScore > 0 {
		percentage := float64(q.Score) / float64(q.TotalScore) * 100
		if percentage >= float64(q.Config.PassingScore) {
			fmt.Println("Congratulations! You passed!")
		} else {
			fmt.Println("Sorry, you did not pass. Keep practicing!")
		}
	}
}
