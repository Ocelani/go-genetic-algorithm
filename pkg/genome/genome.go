package genome

import (
	"log"
	"math/rand"
	"time"
)

type GAConfig struct {
	// Required fields
	NPops        uint
	PopSize      uint
	NGenerations uint
	HofSize      uint
	Model        Model

	// Optional fields
	ParallelEval bool // Whether to evaluate Individuals in parallel or not
	Migrator     Migrator
	MigFrequency uint // Frequency at which migrations occur
	Speciator    Speciator
	Logger       *log.Logger
	Callback     func(ga *GA)
	EarlyStop    func(ga *GA) bool
	RNG          *rand.Rand
}

type GA struct {
	GAConfig

	Populations Populations
	HallOfFame  Individuals
	Age         time.Duration
	Generations uint
}

// Evaluate method returns the fitness of a genome.
// It specifies the problem to deal with and the algorithm only needs it's output.
func Evaluate() (float64, error) {

}

// Mutate method is where we modify an existing solution by tinkering with it's variables.
// The mutate of a solution essentially boils down to the particular problem.
func Mutate(rng *rand.Rand) {

}

// Crossover method combines two individuals.
// The Genome type argument differs from the struct calling the method,
// wich has to be casted into the struct before being able to apply a crossover operator.
// This is due to the fact that Go doesn't provide generics out of the box.
func Crossover(genome Genome, rng *rand.Rand) {

}

// Clone method produces independent copies of the struct to evolve.
// It ensures that pointer fields are not pointing to identical memory addresses.
// This makes the produced clones to not be shallow copies of the cloned genome.
func Clone() Genome {

}
