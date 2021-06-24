package database

import (
	"time"
)

// Review defines the storage form of a beer review
type Review struct {
	ID        int
	BeerID    int
	FirstName string
	LastName  string
	Score     int
	Text      string
	Created   time.Time
}
