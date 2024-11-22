package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Quiz represents the quiz state
type Quiz struct {
	Config     Config
	Questions  []Question
	Score      int
	StartTime  time.Time
	TimeLimit  time.Duration
}

func (q *Quiz) selectQuestions(quizPath string) error {
	rand.Seed(time.Now().UnixNano())
	
	for _, questionGroup := range q.Config.Questions {
		// Randomly select one question from the group
		selectedID := questionGroup[rand.Intn(len(questionGroup))]
		question, err := loadQuestion(quizPath, selectedID)
		if err != nil {
			return err
		}
		q.Questions = append(q.Questions, question)
	}
	
	if q.Config.RandomizeOrder {
		rand.Shuffle(len(q.Questions), func(i, j int) {
			q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
		})
	}
	
	return nil
}

func (q *Quiz) askQuestion(index int) bool {
	question := q.Questions[index]
	fmt.Printf("\nQuestion %d: %s\n", index+1, question.Question)
	
	switch question.Type {
	case "multiple_choice":
		fmt.Println("Options:")
		for i, option := range question.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}
		var answer int
		fmt.Print("Enter your answer (1-", len(question.Options), "): ")
		fmt.Scanf("%d", &answer)
		
		if answer > 0 && answer <= len(question.Options) {
			userAnswer := question.Options[answer-1]
			for _, correctAnswer := range question.Answers {
				if userAnswer == correctAnswer {
					return true
				}
			}
		}
		
	case "fill_in_blank":
		var answer string
		fmt.Print("Your answer: ")
		fmt.Scanf("%s", &answer)
		
		for _, correctAnswer := range question.Answers {
			if answer == correctAnswer {
				return true
			}
		}
		
	case "true_false":
		var answer string
		fmt.Print("Enter true or false: ")
		fmt.Scanf("%s", &answer)
		
		for _, correctAnswer := range question.Answers {
			if answer == correctAnswer {
				return true
			}
		}
	}
	
	return false
}

func (q *Quiz) run() {
	q.StartTime = time.Now()
	q.TimeLimit = time.Duration(q.Config.TimeLimit) * time.Minute
	
	fmt.Printf("\nWelcome to %s!\n", q.Config.Title)
	fmt.Printf("Time limit: %d minutes\n", q.Config.TimeLimit)
	fmt.Printf("Passing score: %d%%\n\n", q.Config.PassingScore)
	
	for i := range q.Questions {
		if time.Since(q.StartTime) > q.TimeLimit {
			fmt.Println("\nTime's up!")
			break
		}
		
		if q.Config.Settings.ShowTimer {
			timeLeft := q.TimeLimit - time.Since(q.StartTime)
			fmt.Printf("\nTime remaining: %v\n", timeLeft.Round(time.Second))
		}
		
		correct := q.askQuestion(i)
		if correct {
			q.Score++
			if q.Config.Settings.ShowFeedbackAfterEach {
				fmt.Println("Correct!")
			}
		} else if q.Config.Settings.ShowFeedbackAfterEach {
			fmt.Println("Incorrect!")
		}
		
		if q.Config.Settings.AllowSkipping {
			var skip string
			fmt.Print("\nPress Enter to continue or 's' to skip: ")
			fmt.Scanf("%s", &skip)
			if skip == "s" {
				continue
			}
		}
	}
	
	percentage := float64(q.Score) / float64(len(q.Questions)) * 100
	fmt.Printf("\nQuiz completed!\nScore: %.1f%%\n", percentage)
	if percentage >= float64(q.Config.PassingScore) {
		fmt.Println("Congratulations! You passed!")
	} else {
		fmt.Println("Sorry, you didn't pass. Try again!")
	}
}
