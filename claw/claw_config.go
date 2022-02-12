package claw

import "time"

type Config struct {
	Id           int32
	Name         string
	Source       string
	TargetType   string
	TargetSource string
	TargetSet    string
	Comments     string
	Deleted      bool
	CreatedTime  time.Time
	UpdatedTime  time.Time
}
