package genome

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

type GAConfig struct {
	// Required fields
	NPops        uint        // The number of populations that will be used: RAND
	PopSize      uint        // The number of individuals inside each population: 400
	NGenerations uint        // For many generations the populations will be evolved: 5000
	HofSize      uint        // How many of the best individuals should be recorded: 40
	Model        eaopt.Model // Determines how to evolve each population of individuals: Steady state model

	// Optional fields
	ParallelEval bool // Whether to evaluate Individuals in parallel or not
	Migrator     eaopt.Migrator
	MigFrequency uint // Frequency at which migrations occur
	Speciator    eaopt.Speciator
	Logger       *log.Logger
	Callback     func(ga *GA)
	EarlyStop    func(ga *GA) bool
	RNG          *rand.Rand
}

type GA struct {
	GAConfig

	Populations eaopt.Populations
	HallOfFame  eaopt.Individuals
	Age         time.Duration
	Generations uint
}

var (
	corpus = strings.Split("abcdefghijklmnopqrstuvwxyz ", "")
	target = strings.Split("software release", "")
)

// Strings is a slice of strings.
type Strings []string

// MakeStrings creates random Strings slices
// by picking random characters from a corpus.
func MakeStrings(rng *rand.Rand) eaopt.Genome {
	return Strings(eaopt.InitUnifString(uint(len(target)), corpus, rng))
}

// Evaluate method returns the fitness of a genome.
// It specifies the problem to deal with and the algorithm only needs it's output.
func (X Strings) Evaluate() (mismatches float64, err error) {
	for i, s := range X {
		if s != target[i] {
			mismatches++
		}
	}
	return
}

// Mutate a Strings slice by replacing it's elements
// by random characters contained in  a corpus.
func (X Strings) Mutate(rng *rand.Rand) {
	eaopt.MutUniformString(X, corpus, 3, rng)
}

// Crossover a Strings slice with another by applying 2-point crossover.
func (X Strings) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXString(X, Y.(Strings), 2, rng)
}

// Clone method produces independent copies of the Strings to evolve.
// Pointer fields are not pointing to identical memory addresses.
// This makes the produced clones to not be shallow copies of the genome.
func (X Strings) Clone() eaopt.Genome {
	var XX = make(Strings, len(X))
	copy(XX, X)
	return XX
}
