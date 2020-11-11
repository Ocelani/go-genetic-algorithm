package pkg

import "math/rand"

// Stakeholder have different importance levels (1 ... 5)
// to the company and different requirements priorities.
type Stakeholder struct {
	Priority     int
	Requirements []*Requirement
}

// NewStakeholder instantiates a new stakeholder.
func NewStakeholder() *Stakeholder {
	i := rand.Intn(5)

	return &Stakeholder{
		Priority:     i,
		Requirements: NewRequirementsList(i),
	}
}
