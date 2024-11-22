package main

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Quoter struct {
	wisdom    int
	humor     string
	power     float64
	knowledge bool
}

func NewQuoter() *Quoter {
	return &Quoter{
		wisdom:    rng.Intn(100),
		humor:     "haha",
		power:     rng.Float64() * 9000,
		knowledge: rng.Intn(2) == 1,
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
	quotes := []string{
		q.GetLifeQuote(),
		q.GetPasswordQuote(),
		q.GetWisdomQuote(),
		q.GetHumorQuote(),
		"May the Force be with you.",
	}
	return quotes[rng.Intn(len(quotes))]
}
