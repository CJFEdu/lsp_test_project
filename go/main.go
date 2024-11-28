package main

import (
	"fmt"
	"os"
	"path/filepath"
	"quiz/quiz_logic"
)

func main() {
	basePath := filepath.Join("..", "quiz")

	quizzes, err := quiz_logic.GetAvailableQuizzes(basePath)

	if err != nil {
		fmt.Printf("Error loading quizzes: %v\n", err)
		os.Exit(1)
	}

	quoter := quiz_logic.NewQuoter()

	for {
		quiz_logic.ShowMenu()
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			quiz_logic.ListQuizzes(quizzes)
		case "2":
			if selectedQuiz := quiz_logic.PromptForQuiz(quizzes); selectedQuiz != nil {
				if err := quiz_logic.StartQuiz(selectedQuiz.Path); err != nil {
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
