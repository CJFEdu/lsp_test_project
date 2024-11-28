package quiz_logic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadConfig(quizPath string) (Config, error) {
	var config Config
	configPath := filepath.Join(quizPath, "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("error reading config: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config: %v", err)
	}

	return config, nil
}

func StartQuiz(quizPath string) error {
	config, err := LoadConfig(quizPath)
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	quiz := Quiz{Config: config}
	err = quiz.selectQuestions(quizPath)
	if err != nil {
		return fmt.Errorf("error loading questions: %v", err)
	}

	quiz.Run()
	return nil
}
