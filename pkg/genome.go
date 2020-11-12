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

var dev = NewDevelopment()

var target = dev.Release

// Evaluate method returns the fitness of a genome.
func (x Release) Evaluate() (mismatches float64, err error) {
	// dev.Requirements = strings.Split("abcdefghijkirthyt", "")
	for i, s := range x {
		if s != dev.Requirements[i] {
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
func (x Release) Crossover(y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXString(x, y.(Release), 2, rng)
}

// Clone method produces independent copies of the Release to evolve.
// Pointer fields are not pointing to identical memory addresses.
// This makes the produced clones to not be shallow copies of the genome.
func (x Release) Clone() eaopt.Genome {
	var xx = make(Release, len(x))
	copy(xx, x)
	return xx
}

// MakeRelease creates random Release strings slices.
func MakeRelease(rng *rand.Rand) eaopt.Genome {
	return Release(eaopt.InitUnifString(uint(len(dev.Requirements)), corpus, rng))
}

// Run executes the algorithm.
func Run() {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	ga.NPops = 400                   // The number of populations that will be used
	ga.PopSize = 100                 // The number of individuals inside each population
	ga.NGenerations = 5000           // For many generations the populations will be evolved
	ga.HofSize = 1                   // How many of the best individuals should be recorded
	ga.Model = eaopt.ModSteadyState{ // Determines how to evolve each population of individuals
		Selector:  eaopt.SelElitism{},
		MutRate:   0.1,
		CrossRate: 0.9,
	}
	ga.ParallelEval = true

	ga.Callback = func(ga *eaopt.GA) { // Add a custom print function to track progress
		var buffer bytes.Buffer

		for _, letter := range ga.HallOfFame[0].Genome.(Release) {
			buffer.WriteString(letter) // display results
		}
		fmt.Printf("%d) Result -> %s (%.0f mismatches)\n",
			ga.Generations, buffer.String(), ga.HallOfFame[0].Fitness,
		)
	}

	// Run the GA
	ga.Minimize(MakeRelease)
}
