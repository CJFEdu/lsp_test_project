package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type QuizInfo struct {
	ID    string
	Title string
	Path  string
}

func getAvailableQuizzes(basePath string) ([]QuizInfo, error) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	var quizzes []QuizInfo
	for _, entry := range entries {
		fmt.Println("Found:", entry.Name())
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "quiz") {
			quizPath := filepath.Join(basePath, entry.Name())
			config, err := loadConfig(quizPath) // loadConfig already joins with config.json
			if err == nil {
				quizzes = append(quizzes, QuizInfo{
					ID:    entry.Name(),
					Title: config.Title,
					Path:  quizPath,
				})
			} else {
				return nil, fmt.Errorf("error loading config for %s: %v", entry.Name(), err)
			}
		}
	}

	// Sort quizzes by ID
	sort.Slice(quizzes, func(i, j int) bool {
		return quizzes[i].ID < quizzes[j].ID
	})

	return quizzes, nil
}

func showMenu() {
	fmt.Println("\n=== Quiz Program Menu ===")
	fmt.Println("1. List Available Quizzes")
	fmt.Println("2. Start a Quiz")
	fmt.Println("3. Exit")
	fmt.Print("\nEnter your choice (1-3): ")
}

func listQuizzes(quizzes []QuizInfo) {
	fmt.Println("\n=== Available Quizzes ===")
	if len(quizzes) == 0 {
		fmt.Println("No quizzes available.")
		return
	}

	fmt.Println("ID\tTitle")
	fmt.Println("--\t-----")
	for _, quiz := range quizzes {
		fmt.Printf("%s\t%s\n", quiz.ID, quiz.Title)
	}
}

func promptForQuiz(quizzes []QuizInfo) *QuizInfo {
	fmt.Print("\nEnter quiz ID (or 'q' to return to menu): ")
	var input string
	fmt.Scanln(&input)

	if input == "q" {
		return nil
	}

	for _, quiz := range quizzes {
		if quiz.ID == input {
			return &quiz
		}
	}

	fmt.Println("Invalid quiz ID.")
	return nil
}
