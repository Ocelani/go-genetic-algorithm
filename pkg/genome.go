package pkg

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

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

// Run executes the algorithm.
func Run() {
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
		RNG:          rand.New(rand.NewSource(42)),
		ParallelEval: true,
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		// Concatenate the elements from the best individual and display the result
		var buffer bytes.Buffer
		for _, letter := range ga.HallOfFame[0].Genome.(Strings) {
			buffer.WriteString(letter)
		}
		fmt.Printf("%d) Result -> %s (%.0f mismatches)\n", ga.Generations, buffer.String(), ga.HallOfFame[0].Fitness)
	}

	// Run the GA
	ga.Minimize(MakeStrings)

	return
}

// 	NPops        uint        // The number of populations that will be used
// 	PopSize      uint        // The number of individuals inside each population
//  NGenerations             // For many generations the populations will be evolved
// 	HofSize      uint        // How many of the best individuals should be recorded
// 	Model        eaopt.Model // Determines how to evolve each population of individuals
