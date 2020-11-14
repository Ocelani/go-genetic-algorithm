package pkg

import (
	"math/rand"
	"strings"
	"time"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

// Requirement is a symbolizes a software requirement.
type Requirement struct {
	Char     string
	Priority int
	StakeID  int
}

// NewRequirement instantiates a new requirements.
func NewRequirement(stk, p int) *Requirement {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := eaopt.InitUniqueString(1, corpus, r)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Requirement{
		StakeID: stk,
		Char:    strings.Join(s, ""),
		Priority: rand.New(
			rand.NewSource(time.Now().UnixNano())).Intn(5),
	}
}

// NewRequirementsList instantiates a new list of requirements.
func NewRequirementsList(stk, p int) []*Requirement {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := eaopt.InitUniqueString(18, corpus, r)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	reqs := []*Requirement{}
	for _, c := range s {
		reqs = append(reqs, &Requirement{
			Char:     string(c),
			Priority: r.Intn(5),
			StakeID:  stk,
		})
	}

	return reqs
}
