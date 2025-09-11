package models

import "time"

type User struct {
	Id          uint64    `json:"id"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Priviledged bool      `json:"priviledged"`
	// Add is admin!! Or something like that

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
	Id     uint64   `json:"id"`
	Source string   `json:"source"`
	Quote  string   `json:"quote"`
	Length QuoteLen `json:"length"`
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
