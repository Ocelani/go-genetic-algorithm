package pkg

import (
	"log"
	"math/rand"
	"time"
)

// Requirements is a list of each requirements
// Req vs. Priority
type Requirements map[string]int

// NewRequirement isntantiates a new requirements.
func NewRequirement(i int) map[string]int {
	s, err := GenerateRandomString(1)
	if err != nil {
		log.Fatal(err)
	}
	return map[string]int{
		s: i + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5),
	}
}

// NewRequirementsList isntantiates a new list of requirementss.
func NewRequirementsList(i int) map[string]int {
	s, err := GenerateRandomString(19)
	if err != nil {
		log.Fatal(err)
	}

	reqs := map[string]int{}

	for _, c := range s {
		reqs = map[string]int{
			string(c): i + rand.New(
				rand.NewSource(time.Now().UnixNano())).Intn(5),
		}
	}
	return reqs
}
