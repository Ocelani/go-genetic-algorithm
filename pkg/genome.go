package pkg

import (
	"math/rand"
	"strings"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

// Release is a slice of strings.
type Release []string

var corpus = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_", "")
var dev = NewDevelopment()

// Evaluate method returns the fitness of a genome.
func (x Release) Evaluate() (mismatches float64, err error) {
	for i, s := range x {
		if s != dev.Target[i] {
			mismatches++
		}
	}
	return
}

// Mutate method sets a Release string slice by replacing it's elements.
func (x Release) Mutate(rng *rand.Rand) {
	eaopt.MutUniformString(x, corpus, 3, rng)
}

// Crossover method sets a Release string slice with another.
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
