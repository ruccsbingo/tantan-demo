package model

import (
// "time"
)

type Relation struct {
	Id          string `json:"id"`
	UserId      int `json:"user_id"`
	OtherUserId int `json:"other_user_id"`
	State       string `json:"state"`
	Type        string `json:"type"`
	// CreateTime time.Time `json:"create_time"`
	// UpdateTime time.Time `json:"update_time"`
}

type Relations []Relation
