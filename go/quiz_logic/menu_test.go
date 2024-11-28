package quiz_logic

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetAvailableQuizzes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "quiz_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test quiz directories and config files
	quizzes := []struct {
		dir   string
		title string
	}{
		{"quiz01", "First Quiz"},
		{"quiz02", "Second Quiz"},
	}

	for _, quiz := range quizzes {
		quizDir := filepath.Join(tempDir, quiz.dir)
		err := os.Mkdir(quizDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create quiz directory: %v", err)
		}

		config := []byte(`{
			"title": "` + quiz.title + `",
			"time_limit": 300,
			"pass_score": 70,
			"questions": []
		}`)

		err = os.WriteFile(filepath.Join(quizDir, "config.json"), config, 0644)
		if err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}
	}

	// Test getting available quizzes
	result, err := GetAvailableQuizzes(tempDir)
	if err != nil {
		t.Errorf("getAvailableQuizzes() error = %v", err)
		return
	}

	if len(result) != len(quizzes) {
		t.Errorf("Expected %d quizzes, got %d", len(quizzes), len(result))
		return
	}

	// Verify quiz information
	for i, quiz := range result {
		if quiz.ID != i+1 {
			t.Errorf("Expected quiz ID %d, got %d", i+1, quiz.ID)
		}
		expectedTitle := quizzes[i].title
		if quiz.Title != expectedTitle {
			t.Errorf("Expected quiz title %q, got %q", expectedTitle, quiz.Title)
		}
	}
}

func TestGetAvailableQuizzes_InvalidDirectory(t *testing.T) {
	_, err := GetAvailableQuizzes("/nonexistent/directory")
	if err == nil {
		t.Error("Expected error for nonexistent directory, got nil")
	}
}

func TestGetAvailableQuizzes_InvalidConfig(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "quiz_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create quiz directory with invalid config
	quizDir := filepath.Join(tempDir, "quiz01")
	err = os.Mkdir(quizDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create quiz directory: %v", err)
	}

	invalidConfig := []byte(`{invalid json}`)
	err = os.WriteFile(filepath.Join(quizDir, "config.json"), invalidConfig, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	_, err = GetAvailableQuizzes(tempDir)
	if err == nil {
		t.Error("Expected error for invalid config, got nil")
	}
}
