package model

import "time"

type Student struct {
	ID             int64
	Name           string
	Ruby           string
	StudentNumber  int
	DateOfBirth    time.Time
	Address        string
	ExpirationDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
