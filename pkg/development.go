package pkg

import (
	"container/heap"
)

// Development represents a software project variables.
type Development struct {
	Requirements []*Requirement // 19 requirements
	Stakeholders []*Stakeholder // 5 stakeholders
	Releases     int            // 5 releases
	Resources    int            // 3 different resources
}

// Resources are all sort of goods available
// for the completion of the project activities.
// groupID vs. amount (amount < effort)
type Resources map[int]int

// NewDevelopment instantiates a new Development type.
func NewDevelopment() (dev *Development) {
	dev.setStakeholders(5)
	dev.setRequirements()
	dev.Resources = 3
	dev.Releases = 0
	return
}

func (dev *Development) setStakeholders(n int) {
	for i := 0; i < n; i++ {
		dev.Stakeholders = append(dev.Stakeholders, NewStakeholder())
	}
	return
}

func (dev *Development) setRequirements() {
	var (
		queue PriorityQueue
		i     int
	)

	for _, stk := range dev.Stakeholders {
		for _, req := range stk.Requirements {
			queue[i] = &Item{
				Requirement: Requirement{
					String: req.String, PriorityRisk: req.PriorityRisk,
				},
				index: i,
			}
			i++
		}
	}
	heap.Init(&queue)

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		dev.Requirements = append(dev.Requirements, &item.Requirement)
	}

	return
}
