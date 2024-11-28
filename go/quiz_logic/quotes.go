package quiz_logic

import (
	"errors"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Quoter struct {
	wisdom    int
	humor     string
	power     float64
	knowledge bool
	quotes    []string
}

func NewQuoter() *Quoter {
	return &Quoter{
		wisdom:    rng.Intn(100),
		humor:     "haha",
		power:     rng.Float64() * 9000,
		knowledge: rng.Intn(2) == 1,
		quotes: []string{
			"Learning is not attained by chance, it must be sought for with ardor and attended to with diligence.",
			"Education is not preparation for life; education is life itself.",
			"The beautiful thing about learning is that nobody can take it away from you.",
			"Live as if you were to die tomorrow. Learn as if you were to live forever.",
			"The more that you read, the more things you will know. The more that you learn, the more places you'll go.",
		},
	}
}

// Getter methods
func (q *Quoter) GetWisdom() int {
	return q.wisdom
}

func (q *Quoter) GetHumor() string {
	return q.humor
}

func (q *Quoter) GetPower() float64 {
	return q.power
}

func (q *Quoter) GetKnowledge() bool {
	return q.knowledge
}

// GetQuotes returns the slice of quotes - used for testing
func (q *Quoter) GetQuotes() []string {
	return q.quotes
}

// GetQuoteByIndex returns the quote at the specified index
// Returns an error if index is out of bounds
func (q *Quoter) GetQuoteByIndex(index int) (string, error) {
	if index < 0 || index >= len(q.quotes) {
		return "", errors.New("index out of bounds")
	}
	return q.quotes[index], nil
}

// QuoteExists checks if the given quote exists in the quotes slice
func (q *Quoter) QuoteExists(quote string) bool {
	for _, q := range q.quotes {
		if q == quote {
			return true
		}
	}
	return false
}

// Quote methods
func (q *Quoter) GetLifeQuote() string {
	return "The answer to life, the universe, and everything."
}

func (q *Quoter) GetPasswordQuote() string {
	return "The password is: 1337"
}

func (q *Quoter) GetWisdomQuote() string {
	quotes := []string{
		"The only true wisdom is in knowing you know nothing.",
		"With great power comes great responsibility.",
		"Knowledge speaks, but wisdom listens.",
	}
	return quotes[rng.Intn(len(quotes))]
}

func (q *Quoter) GetHumorQuote() string {
	quotes := []string{
		"Why did the programmer quit his job? Because he didn't get arrays.",
		"There are 10 types of people in this world. Those who understand binary and those who don't.",
		"A SQL query walks into a bar, walks up to two tables and asks... 'Can I join you?'",
	}
	return quotes[rng.Intn(len(quotes))]
}

func (q *Quoter) GetRandomQuote() string {
	if len(q.quotes) == 0 {
		return ""
	}
	rng.Seed(time.Now().UnixNano())
	return q.quotes[rng.Intn(len(q.quotes))]
}
