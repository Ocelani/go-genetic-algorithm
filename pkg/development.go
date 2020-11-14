package pkg

import (
	"container/heap"
	"math/rand"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
	"github.com/Ocelani/go-genetic-algorithm/pkg/queue"
)

// Development represents a software project variables.
type Development struct {
	Stakeholders []*Stakeholder // 5 stakeholders
	Requirements []*Requirement // 19 requirements
	Release      Release        // 5 releases
	Target       []string       // 5 releases
	// Resources    int            // 3 different resources
}

// Resources are all sort of goods available
// for the completion of the project activities.
type Resources map[int]int

// NewDevelopment instantiates a new Development type.
func NewDevelopment() *Development {
	dev := &Development{}
	dev.setStakeholders(5)
	dev.setProjectRequirements()
	return dev
}

// MakeRelease method creates random Release string slices.
func (dev *Development) MakeRelease(rng *rand.Rand) eaopt.Genome {
	dev.Release = Release(eaopt.InitUnifString(uint(len(dev.Release)), corpus, rng))
	return dev.Release
}

func (dev *Development) setStakeholders(n int) {
	for i := 0; i < n; i++ {
		dev.Stakeholders = append(dev.Stakeholders, NewStakeholder(i+1))
	}
}

func (dev *Development) setProjectRequirements() {
	i := 0
	pq := queue.PriorityQueue{}

	for _, stk := range dev.Stakeholders {
		for _, r := range stk.Requirements {
			pq = append(pq,
				&queue.Item{
					Index:    i,
					Value:    r.Char,
					Priority: r.Priority + stk.Priority,
					StakeID:  stk.ID,
				},
			)
			i++
		}
	}
	heap.Init(&pq)

	dev.Target = []string{}
	for len(dev.Requirements) < 19 {
		item := heap.Pop(&pq).(*queue.Item)
		dev.Target = append(dev.Target, item.Value)
		dev.Requirements = append(dev.Requirements, &Requirement{
			Char:     item.Value,
			Priority: item.Priority,
			StakeID:  item.StakeID,
		})
	}
	dev.Release = make([]string, len(dev.Requirements))
}
