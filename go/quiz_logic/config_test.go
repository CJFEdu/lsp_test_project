package quiz_logic

import (
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	baseDir := "../../quiz"

	// Test cases
	tests := []struct {
		name              string
		path              string
		wantErr           bool
		expectedTitle     string
		expectedTime      int
		expectedScore     int
		expectedRandom    bool
		expectedQuestions int
		expectedSettings  struct {
			ShowFeedbackAfterEach bool
			AllowSkipping         bool
			ShowTimer             bool
		}
		expectedQuestionSet [][]string
	}{
		{
			name:              "Basic Quiz Config",
			path:              filepath.Join(baseDir, "test01"),
			wantErr:           false,
			expectedTitle:     "Basic Knowledge Test",
			expectedTime:      5,
			expectedScore:     60,
			expectedRandom:    false,
			expectedQuestions: 4,
			expectedSettings: struct {
				ShowFeedbackAfterEach bool
				AllowSkipping         bool
				ShowTimer             bool
			}{
				ShowFeedbackAfterEach: true,
				AllowSkipping:         false,
				ShowTimer:             true,
			},
			expectedQuestionSet: [][]string{
				{"question001"},
				{"question002"},
				{"question003"},
				{"question004"},
			},
		},
		{
			name:              "Alternative Questions Config",
			path:              filepath.Join(baseDir, "test02"),
			wantErr:           false,
			expectedTitle:     "Alternative Questions Test",
			expectedTime:      15,
			expectedScore:     80,
			expectedRandom:    true,
			expectedQuestions: 4,
			expectedSettings: struct {
				ShowFeedbackAfterEach bool
				AllowSkipping         bool
				ShowTimer             bool
			}{
				ShowFeedbackAfterEach: false,
				AllowSkipping:         true,
				ShowTimer:             false,
			},
			expectedQuestionSet: [][]string{
				{"question001", "question002"},
				{"question003", "question004"},
				{"question002", "question005"},
				{"question001", "question003"},
			},
		},
		{
			name:              "Quick Quiz Config",
			path:              filepath.Join(baseDir, "test03"),
			wantErr:           false,
			expectedTitle:     "Quick Quiz",
			expectedTime:      0,
			expectedScore:     0,
			expectedRandom:    false,
			expectedQuestions: 2,
			expectedSettings: struct {
				ShowFeedbackAfterEach bool
				AllowSkipping         bool
				ShowTimer             bool
			}{
				ShowFeedbackAfterEach: true,
				AllowSkipping:         false,
				ShowTimer:             false,
			},
			expectedQuestionSet: [][]string{
				{"question001"},
				{"question005"},
			},
		},
		{
			name:    "Non-existent directory",
			path:    filepath.Join(baseDir, "nonexistent"),
			wantErr: true,
		},
		{
			name:    "Invalid parent directory",
			path:    filepath.Join(baseDir, "nonexistent", "test01"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := LoadConfig(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify basic config contents
				if config.Title != tt.expectedTitle {
					t.Errorf("Expected title %q, got %q", tt.expectedTitle, config.Title)
				}
				if config.TimeLimit != tt.expectedTime {
					t.Errorf("Expected time limit %d, got %d", tt.expectedTime, config.TimeLimit)
				}
				if config.PassingScore != tt.expectedScore {
					t.Errorf("Expected passing score %d, got %d", tt.expectedScore, config.PassingScore)
				}
				if config.RandomizeOrder != tt.expectedRandom {
					t.Errorf("Expected randomize order %v, got %v", tt.expectedRandom, config.RandomizeOrder)
				}

				// Verify settings
				if config.Settings.ShowFeedbackAfterEach != tt.expectedSettings.ShowFeedbackAfterEach {
					t.Errorf("Expected ShowFeedbackAfterEach %v, got %v",
						tt.expectedSettings.ShowFeedbackAfterEach,
						config.Settings.ShowFeedbackAfterEach)
				}
				if config.Settings.AllowSkipping != tt.expectedSettings.AllowSkipping {
					t.Errorf("Expected AllowSkipping %v, got %v",
						tt.expectedSettings.AllowSkipping,
						config.Settings.AllowSkipping)
				}
				if config.Settings.ShowTimer != tt.expectedSettings.ShowTimer {
					t.Errorf("Expected ShowTimer %v, got %v",
						tt.expectedSettings.ShowTimer,
						config.Settings.ShowTimer)
				}

				// Verify questions array
				if len(config.Questions) != tt.expectedQuestions {
					t.Errorf("Expected %d questions, got %d", tt.expectedQuestions, len(config.Questions))
				}

				// Verify question sets structure
				for i, expectedSet := range tt.expectedQuestionSet {
					if i >= len(config.Questions) {
						t.Errorf("Missing question set at index %d", i)
						continue
					}
					actualSet := config.Questions[i]
					if len(actualSet) != len(expectedSet) {
						t.Errorf("Question set %d: expected %d alternatives, got %d",
							i, len(expectedSet), len(actualSet))
						continue
					}
					for j, expectedQuestion := range expectedSet {
						if actualSet[j] != expectedQuestion {
							t.Errorf("Question set %d, alternative %d: expected %q, got %q",
								i, j, expectedQuestion, actualSet[j])
						}
					}
				}
			}
		})
	}
}
