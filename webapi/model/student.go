package model

import (
	"time"
)

type Student struct {
	ID             uint64
	Name           string
	Ruby           string
	StudentNumber  uint32
	DateOfBirth    time.Time
	Address        string
	ExpirationDate time.Time
}
