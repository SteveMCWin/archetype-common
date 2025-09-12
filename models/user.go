package models

import (
	"strings"
	"time"
)

type User struct {
	Id          uint64    `json:"id"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Priviledged bool      `json:"priviledged"`

	// stats
	TestsStarted   uint    `json:"testsStarted"`
	TestsCompleted uint    `json:"testsCompleted"`
	AllTimeAvgWPM  float32 `json:"allTimeAvgWPM"`
	AllTimeAvgACC  float32 `json:"allTimeAvgACC"`
}

type QuoteLen int

const (
	QUOTE_SHORT QuoteLen = iota
	QUOTE_MEDIUM
	QUOTE_LARGE
	QUOTE_EXTRA_LARGE
	QUOTE_ANY_SIZE
)

type Quote struct {
	Id         uint64   `json:"id"`
	Source     string   `json:"source"`
	Quote      string   `json:"quote"`
	Length     QuoteLen `json:"length"`
	NumOfStars int      `json:"num_of_stars"`
}

func CalcQuoteLen(quote_text string) QuoteLen {
	num_of_words := len(strings.Split(quote_text, " "))
	switch {
	case num_of_words < 21:
		return QUOTE_SHORT
	case num_of_words < 46:
		return QUOTE_MEDIUM
	case num_of_words < 101:
		return QUOTE_LARGE
	default:
		return QUOTE_EXTRA_LARGE
	}
}

type QuoteRequest struct {
	Length QuoteLen `json:"length"`
}

type WordsRequest struct {
	Length      uint `json:"length"`
	TimeLimited bool `json:"time_limited"`
	Punctuation bool `json:"punctuation"`
	Numbers     bool `json:"numbers"`
}
