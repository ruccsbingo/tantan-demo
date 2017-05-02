package model

import (
// "time"
)

type Account struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// CreateTime time.Time `json:"create_time"`
	// UpdateTime time.Time `json:"update_time"`
}

type Accounts []Account
