package pkg

import (
	"log"
	"math/rand"
	"time"
)

// Requirement is a list of each requirement
//  with a risk value (0 ... 5).
// Req vs. PriorityRisk
type Requirement struct {
	String       string
	PriorityRisk int
}

// NewRequirement isntantiates a new requirement.
func NewRequirement(i int) *Requirement {
	s, err := GenerateRandomString(1)
	if err != nil {
		log.Fatal(err)
	}

	return &Requirement{
		String:       s,
		PriorityRisk: i + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5),
	}
}

// NewRequirementsList isntantiates a new list of requirements.
func NewRequirementsList(i int) []*Requirement {
	s, err := GenerateRandomString(19)
	if err != nil {
		log.Fatal(err)
	}

	var reqs []*Requirement

	for _, c := range s {
		reqs = append(reqs, &Requirement{
			String:       string(c),
			PriorityRisk: i + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5),
		})
	}

	return reqs
}
