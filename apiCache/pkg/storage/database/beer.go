package database

import (
	"time"
)

// Beer defines the properties of a beer to be listed
type Beer struct {
	ID        string
	Name      string
	Brewery   string
	Abv       float32
	ShortDesc string
	Created   time.Time
}
