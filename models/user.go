package models

import "time"

type User struct {
	Id          uint64    `json:"id"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`

	// stats
	TestsStarted   uint    `json:"testsStarted"`
	TestsCompleted uint    `json:"testsCompleted"`
	AllTimeAvgWPM  float32 `json:"allTimeAvgWPM"`
	AllTimeAvgACC  float32 `json:"allTimeAvgACC"`
}

type QuoteLen string

const (
	QUOTE_SHORT       QuoteLen = "s"
	QUOTE_MEDIUM      QuoteLen = "m"
	QUOTE_LARGE       QuoteLen = "l"
	QUOTE_EXTRA_LARGE QuoteLen = "xl"
)

type Quote struct {
	Id     int      `json:"id"`
	Source string   `json:"source"`
	Quote  string   `json:"quote"`
	Length QuoteLen `json:"length"`
}
