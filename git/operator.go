package git

import "time"

// Operator ...
type Operator struct {
	Name  string
	Email string
	Time  time.Time
	Zone  string
}
