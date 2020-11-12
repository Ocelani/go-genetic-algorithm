package pkg

// Stakeholder have different importance levels (1 ... 5)
// to the company and different requirements priorities.
type Stakeholder struct {
	Priority     int
	Requirements map[string]int
}

// NewStakeholder instantiates a new stakeholder.
func NewStakeholder() Stakeholder {
	i := 4

	return Stakeholder{
		Priority:     i,
		Requirements: NewRequirementsList(i),
	}
}
