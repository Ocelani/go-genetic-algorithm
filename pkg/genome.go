package pkg

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

// TODO: re-factor

// Release is a slice of Release.
type Release []string

var corpus = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_", "")

// Evaluate method returns the fitness of a genome.
func (x Release) Evaluate() (mismatches float64, err error) {
	for i, s := range x {
		if s != x[i] {
			mismatches++
		}
	}
	return
}

// Mutate method sets a Release string slice by replacing it's elements.
func (x Release) Mutate(rng *rand.Rand) {
	eaopt.MutUniformString(x, corpus, 3, rng)
}

// Crossover method sets a Release string slice with another by applying 2-point crossover.
func (x Release) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXString(x, Y.(Release), 2, rng)
}

// Clone method produces independent copies of the Release to evolve.
// Pointer fields are not pointing to identical memory addresses.
// This makes the produced clones to not be shallow copies of the genome.
func (x Release) Clone() eaopt.Genome {
	var xx = make(Release, len(x))
	copy(xx, x)
	return xx
}

// Run executes the algorithm.
func Run() {
	dev := NewDevelopment()

	c := &eaopt.GAConfig{
		NPops:        400,  // The number of populations that will be used
		PopSize:      100,  // The number of individuals inside each population
		NGenerations: 5000, // For many generations the populations will be evolved
		HofSize:      1,    // How many of the best individuals should be recorded
		Model: eaopt.ModSteadyState{ // Determines how to evolve each population of individuals
			Selector:  eaopt.SelElitism{},
			MutRate:   0.1,
			CrossRate: 0.9,
		},
		// RNG:          rand.New(rand.NewSource(42)),
		ParallelEval: false,
		// EarlyStop:    func(ga *eaopt.GA) bool { if ga.HallOfFame[] },
	}
	ga, err := c.NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		var buffer bytes.Buffer
		// Concatenate the elements from the best individual and display the result
		for _, letter := range ga.HallOfFame[0].Genome.(Release) {
			buffer.WriteString(letter)
		}
		fmt.Printf("%d) Result -> %s (%.0f mismatches)\n",
			ga.Generations, buffer.String(), ga.HallOfFame[0].Fitness,
		)
	}

	// Run the GA
	ga.Minimize(dev.MakeRelease)
}
