package pkg

import (
	"container/heap"
	"math/rand"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

// Development represents a software project variables.
type Development struct {
	Stakeholders []*Stakeholder // 5 stakeholders
	Requirements map[string]int // 19 requirements
	Resources    int            // 3 different resources
	Release      Release        // 5 releases
	Target       []string       // 5 releases
}

// Resources are all sort of goods available
// for the completion of the project activities.
type Resources map[int]int

// NewDevelopment instantiates a new Development type.
func NewDevelopment() *Development {
	dev := &Development{Requirements: map[string]int{}}
	dev.setStakeholders(5)
	dev.setProjectRequirements()
	dev.Resources = 3
	return dev
}

// MakeRelease method creates random Release string slices.
func (dev *Development) MakeRelease(rng *rand.Rand) eaopt.Genome {
	dev.Release = Release(eaopt.InitUnifString(uint(len(dev.Release)), corpus, rng))
	return dev.Release
}

func (dev *Development) setStakeholders(n int) {
	for i := 0; i < n; i++ {
		dev.Stakeholders = append(dev.Stakeholders, NewStakeholder())
	}
}

func (dev *Development) setProjectRequirements() {
	i := 0
	queue := PriorityQueue{}

	for _, stk := range dev.Stakeholders {
		for v, p := range stk.Requirements {
			queue = append(queue,
				&Item{
					Index:    i,
					Value:    v,
					Priority: p,
				},
			)
			i++
		}
	}
	heap.Init(&queue)

	dev.Target = []string{}
	for len(dev.Requirements) < 20 {
		item := heap.Pop(&queue).(*Item)
		dev.Requirements[item.Value] = item.Priority
		dev.Target = append(dev.Target, item.Value)
	}
	dev.Release = make([]string, len(dev.Requirements))
}
