package pkg

import (
	"math/rand"
	"time"
)

// Stakeholder have different importance levels (1 ... 5)
// to the company and different requirements priorities.
type Stakeholder struct {
	ID           int
	Priority     int
	Requirements []*Requirement
}

// NewStakeholder instantiates a new stakeholder.
func NewStakeholder(i int) *Stakeholder {
	p := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	).Intn(5)

	return &Stakeholder{
		ID:           i,
		Priority:     p,
		Requirements: NewRequirementsList(i, p),
	}
}
