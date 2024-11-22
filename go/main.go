package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func loadConfig(quizPath string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filepath.Join(quizPath, "config.json"))
	if err != nil {
		return config, fmt.Errorf("error reading config: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config: %v", err)
	}

	return config, nil
}

func loadQuestion(quizPath, questionID string) (Question, error) {
	var question Question
	data, err := os.ReadFile(filepath.Join(quizPath, questionID+".json"))
	if err != nil {
		return question, fmt.Errorf("error reading question %s: %v", questionID, err)
	}

	err = json.Unmarshal(data, &question)
	if err != nil {
		return question, fmt.Errorf("error parsing question %s: %v", questionID, err)
	}

	return question, nil
}

func startQuiz(quizPath string) error {
	config, err := loadConfig(filepath.Join(quizPath, "config.json"))
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	quiz := Quiz{Config: config}
	err = quiz.selectQuestions(quizPath)
	if err != nil {
		return fmt.Errorf("error loading questions: %v", err)
	}

	quiz.run()
	return nil
}

func main() {
	basePath := filepath.Join("..", "quiz")

	quizzes, err := getAvailableQuizzes(basePath)

	if err != nil {
		fmt.Printf("Error loading quizzes: %v\n", err)
		os.Exit(1)
	}

	quoter := NewQuoter()

	for {
		showMenu()
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			listQuizzes(quizzes)
		case "2":
			if selectedQuiz := promptForQuiz(quizzes); selectedQuiz != nil {
				if err := startQuiz(selectedQuiz.Path); err != nil {
					fmt.Printf("Error running quiz: %v\n", err)
				}
			}
		case "3":
			fmt.Println("Goodbye!")
			return
		case "42":
			fmt.Println(quoter.GetLifeQuote())
		case "1337":
			fmt.Println(quoter.GetPasswordQuote())
		case "1234":
			fmt.Println(quoter.GetWisdomQuote())
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
