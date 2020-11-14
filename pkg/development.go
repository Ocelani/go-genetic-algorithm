package pkg

import (
	"container/heap"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

// Development represents a software project variables.
type Development struct {
	Stakeholders []*Stakeholder // 5 stakeholders
	Requirements map[string]int // 19 requirements
	Release      Release        // 5 releases
	Target       []string       // 5 releases
	// Resources    int         // 3 different resources
}

// Resources are all sort of goods available
// for the completion of the project activities.
type Resources map[int]int

// NewDevelopment instantiates a new Development type.
func NewDevelopment() *Development {
	dev := &Development{Requirements: map[string]int{}}
	dev.setStakeholders(5)
	dev.setProjectRequirements()

	go func() {
		var s, p string
		for sr, pr := range dev.Requirements {
			s = s + " " + sr
			p = p + " " + strconv.Itoa(pr)
		}
		fmt.Printf(`
    # REQUIREMENTS
      requirements: %v
      priorities:   %v
      `, s, p)
		fmt.Println()
	}()

	return dev
}

// MakeRelease method creates random Release string slices.
func (dev *Development) MakeRelease(rng *rand.Rand) eaopt.Genome {
	dev.Release = Release(eaopt.InitUnifString(uint(len(dev.Release)), corpus, rng))
	return dev.Release
}

func (dev *Development) setStakeholders(n int) {
	fmt.Printf(`
  -----------------------------
  STAKEHOLDERS`)

	for i := 0; i < n; i++ {
		stk := NewStakeholder(i + 1)
		dev.Stakeholders = append(dev.Stakeholders, stk)

		var s, p string
		for sr, pr := range stk.Requirements {
			s = s + " " + sr
			p = p + " " + strconv.Itoa(pr)
		}
		fmt.Printf(`
      id: %v
      requirements: %v
      priorities:   %v
      _____________________________
      `, i+1, s, p)
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
