package pkg

import (
	"container/heap"
	"strings"
)

// Development represents a software project variables.
type Development struct {
	Stakeholders []Stakeholder // 5 stakeholders
	Requirements []string      // 19 requirements
	Resources    int           // 3 different resources
	// Release      []string       // 5 releases
}

// Resources are all sort of goods available
// for the completion of the project activities.
type Resources map[int]int

// NewDevelopment instantiates a new Development type.
func NewDevelopment() Development {
	dev := Development{
		Requirements: strings.Split("injOhgdTlnDSArwLpmn", ""),
		// Stakeholders: setStakeholders(5),
	}

	return dev
}

// // MakeRelease method creates random Release string slices.
// func (dev *Development) MakeRelease(rng *rand.Rand) eaopt.Genome {
// 	var release Release
// 	dev.Release = Release{}

// 	for _, d := range dev.Requirements {
// 		release = append(release, d)
// 	}
// 	return Release(eaopt.InitUnifString(uint(len(release)), corpus, rng))
// }

func setStakeholders(n int) []Stakeholder {
	stks := []Stakeholder{}
	for i := 0; i < n; i++ {
		stks = append(stks, NewStakeholder())
	}
	return stks
}

func setProjectRequirements() []string {
	stks := []Stakeholder{}
	for i := 0; i < 5; i++ {
		stks = append(stks, NewStakeholder())
	}

	var queue PriorityQueue

	i := 0
	for _, stk := range stks {
		for v, p := range stk.Requirements {
			queue[i] = &Item{
				Index:    i,
				Value:    v,
				Priority: p,
			}
			i++
		}
	}
	heap.Init(&queue)

	reqs := []string{}
	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		reqs = append(reqs, item.Value)
	}

	return reqs
}
