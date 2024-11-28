package quiz_logic

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type QuizInfo struct {
	ID    int
	Title string
	Path  string
}

func GetAvailableQuizzes(basePath string) ([]QuizInfo, error) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	var quizzes []QuizInfo
	quizID := 1
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "quiz") {
			quizPath := filepath.Join(basePath, entry.Name())
			config, err := LoadConfig(quizPath)
			if err == nil {
				quizzes = append(quizzes, QuizInfo{
					ID:    quizID,
					Title: config.Title,
					Path:  quizPath,
				})
				quizID++
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

func ShowMenu() {
	fmt.Println("\n=== Quiz Program Menu ===")
	fmt.Println("1. List Available Quizzes")
	fmt.Println("2. Start a Quiz")
	fmt.Println("3. Exit")
	fmt.Print("\nEnter your choice (1-3): ")
}

func ListQuizzes(quizzes []QuizInfo) {
	fmt.Println("\n=== Available Quizzes ===")
	if len(quizzes) == 0 {
		fmt.Println("No quizzes available.")
		return
	}

	fmt.Println("ID\tTitle")
	fmt.Println("--\t-----")
	for _, quiz := range quizzes {
		fmt.Printf("%d\t%s\n", quiz.ID, quiz.Title)
	}
}

func PromptForQuiz(quizzes []QuizInfo) *QuizInfo {
	fmt.Print("\nEnter quiz number (or 0 to return to menu): ")
	var input int
	_, err := fmt.Scanln(&input)

	if err != nil || input == 0 {
		return nil
	}

	for _, quiz := range quizzes {
		if quiz.ID == input {
			return &quiz
		}
	}

	fmt.Println("Invalid quiz number.")
	return nil
}
