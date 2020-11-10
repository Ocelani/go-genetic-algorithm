package pkg

import "strings"

// Resources are all sort of goods available
// for the completion of the project activities.
// groupID vs. amount (amount < effort)
type Resources map[int]int

// Requirements is a list of each requirement
//  with a risk value (0 ... 5).
// Req vs. PriorityRisk
type Requirements map[string]int

// Stakeholder have different importance levels (1 ... 5)
// to the company and different requirements priorities.
type Stakeholder struct {
	Priority     int
	Requirements Requirements
}

// Development represents a software project variables.
type Development struct {
	Requirements Requirements  // 19 requirements
	Resources    Resources     // 3 different resources
	Stakeholders []Stakeholder // 5 stakeholders
	Releases     int           // 5 releases
}

var (
	corpus = strings.Split("abcdefghijklmnopqrstuvwxyz ", "")
	target = strings.Split("software release", "")
)

// Strings is a slice of strings.
type Strings []string
