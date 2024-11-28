package quiz_logic

import (
	"testing"
)

func TestNewQuoter(t *testing.T) {
	quoter := NewQuoter()

	if quoter == nil {
		t.Error("NewQuoter() returned nil")
	}

	// Test getter methods
	if quoter.GetWisdom() < 0 || quoter.GetWisdom() > 100 {
		t.Error("Wisdom value out of expected range")
	}

	if quoter.GetHumor() != "haha" {
		t.Errorf("Expected humor to be 'haha', got %q", quoter.GetHumor())
	}

	if quoter.GetPower() < 0 || quoter.GetPower() > 9000 {
		t.Error("Power value out of expected range")
	}

	// Knowledge is boolean, no range check needed
	_ = quoter.GetKnowledge()
}

func TestQuoter_GetRandomQuote(t *testing.T) {
	quoter := NewQuoter()

	// Test multiple times to ensure randomness
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		quote := quoter.GetRandomQuote()
		if quote == "" {
			t.Error("GetRandomQuote() returned empty string")
		}
		seen[quote] = true
	}

	// Check that we got at least a few different quotes
	if len(seen) < 3 {
		t.Error("GetRandomQuote() does not seem to be random enough")
	}
}

func TestQuoter_GetQuoteByIndex(t *testing.T) {
	quoter := NewQuoter()

	tests := []struct {
		name    string
		index   int
		wantErr bool
	}{
		{
			name:    "Valid index",
			index:   0,
			wantErr: false,
		},
		{
			name:    "Last valid index",
			index:   4, // We know there are 5 quotes
			wantErr: false,
		},
		{
			name:    "Negative index",
			index:   -1,
			wantErr: true,
		},
		{
			name:    "Out of bounds index",
			index:   5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quote, err := quoter.GetQuoteByIndex(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuoteByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && quote == "" {
				t.Error("GetQuoteByIndex() returned empty string for valid index")
			}
		})
	}
}

func TestQuoter_SpecialQuotes(t *testing.T) {
	quoter := NewQuoter()

	// Test GetLifeQuote
	lifeQuote := quoter.GetLifeQuote()
	if lifeQuote == "" {
		t.Error("GetLifeQuote() returned empty string")
	}

	// Test GetPasswordQuote
	passQuote := quoter.GetPasswordQuote()
	if passQuote == "" {
		t.Error("GetPasswordQuote() returned empty string")
	}

	// Test GetWisdomQuote
	wisdomQuote := quoter.GetWisdomQuote()
	if wisdomQuote == "" {
		t.Error("GetWisdomQuote() returned empty string")
	}

	// Test GetHumorQuote
	humorQuote := quoter.GetHumorQuote()
	if humorQuote == "" {
		t.Error("GetHumorQuote() returned empty string")
	}
}

func TestQuoter_QuoteExists(t *testing.T) {
	quoter := NewQuoter()

	tests := []struct {
		name  string
		quote string
		want  bool
	}{
		{
			name:  "Existing quote",
			quote: "Live as if you were to die tomorrow. Learn as if you were to live forever.",
			want:  true,
		},
		{
			name:  "Non-existing quote",
			quote: "This quote does not exist",
			want:  false,
		},
		{
			name:  "Empty quote",
			quote: "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quoter.QuoteExists(tt.quote); got != tt.want {
				t.Errorf("QuoteExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
