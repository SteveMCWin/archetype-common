package models

import "time"

type User struct {
	Id          uint64
	UserName    string
	Password    string
	Email       string
	DateCreated time.Time

	// stats
	TestsStarted   uint
	TestsCompleted uint
	AllTimeAvgWPM  float32
	AllTimeAvgACC  float32
}
