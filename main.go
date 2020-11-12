package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// MutationRate is the rate of mutation
var MutationRate = 0.1

// PopSize is the size of the population
var PopSize = 500

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	target := []byte("SoftwareReleaseDay!")
	population := createPopulation(target)

	found := false
	generation := 0
	bestOrganism := getBest(population)
	maxFit := bestOrganism.Fitness
	maxDNA := string(bestOrganism.DNA)
	maxG := 0

	for !found && generation < 5000 {
		generation++
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestOrganism.DNA), bestOrganism.Fitness)

		if bytes.Compare(bestOrganism.DNA, target) == 0 {
			found = true
			elapsed := time.Since(start)
			fmt.Printf("\nTime taken: %s\n", elapsed)
			break
		}

		if bestOrganism.Fitness > maxFit {
			maxFit = bestOrganism.Fitness
			maxDNA = string(bestOrganism.DNA)
			maxG = generation
		}

		bestOrganism = getBest(population)
		pool := createPool(population, target, bestOrganism.Fitness)
		population = naturalSelection(pool, population, target)
	}

	if !found {
		fmt.Printf("\n BEST: generation: %d | %s | fitness: %2f \n", maxG, maxDNA, maxFit)
	}
}

// Organism for this genetic algorithm
type Organism struct {
	DNA     []byte
	Fitness float64
}

// creates a Organism
func createOrganism(target []byte) (organism Organism) {
	ba := make([]byte, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = byte(rand.Intn(95) + 32)
	}
	organism = Organism{
		DNA:     ba,
		Fitness: 0,
	}
	organism.calcFitness(target)
	return
}

// creates the initial population
func createPopulation(target []byte) (population []Organism) {
	population = make([]Organism, PopSize)
	for i := 0; i < PopSize; i++ {
		population[i] = createOrganism(target)
	}
	return
}

// calculates the fitness of the Organism
func (d *Organism) calcFitness(target []byte) {
	score := 0
	for i := 0; i < len(d.DNA); i++ {
		if d.DNA[i] == target[i] {
			score++
		}
	}
	d.Fitness = float64(score) / float64(len(d.DNA))
	return
}

// create the breeding pool that creates the next generation
func createPool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// perform natural selection to create the next generation
func naturalSelection(pool []Organism, population []Organism, target []byte) []Organism {
	next := make([]Organism, len(population))

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := crossover(a, b)
		child.mutate()
		child.calcFitness(target)

		next[i] = child
	}
	return next
}

// crosses over 2 Organisms
func crossover(d1 Organism, d2 Organism) Organism {
	child := Organism{
		DNA:     make([]byte, len(d1.DNA)),
		Fitness: 0,
	}
	mid := rand.Intn(len(d1.DNA))
	for i := 0; i < len(d1.DNA); i++ {
		if i > mid {
			child.DNA[i] = d1.DNA[i]
		} else {
			child.DNA[i] = d2.DNA[i]
		}

	}
	return child
}

// mutate the Organism
func (d *Organism) mutate() {
	for i := 0; i < len(d.DNA); i++ {
		if rand.Float64() < MutationRate {
			d.DNA[i] = byte(rand.Intn(95) + 32)
		}
	}
}

// Get the best organism
func getBest(population []Organism) Organism {
	best := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > best {
			index = i
			best = population[i].Fitness
		}
	}
	return population[index]
}

// // func main() {
// // 	pkg.Run()
// // }

// var (
// 	dev    = pkg.NewDevelopment()
// 	corpus = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_", "")
// )

// // Strings is a slice of strings.
// type Strings []string

// // Evaluate a Strings slice by counting the number of mismatches between itself
// // and the target string.
// func (X Strings) Evaluate() (mismatches float64, err error) {
// 	for i, s := range X {
// 		if s != dev.Requirements[i] {
// 			mismatches++
// 		}
// 	}
// 	return
// }

// // Mutate a Strings slice by replacing it's elements by random characters
// // contained in  a corpus.
// func (X Strings) Mutate(rng *rand.Rand) {
// 	eaopt.MutUniformString(X, corpus, 3, rng)
// }

// // Crossover a Strings slice with another by applying 2-point crossover.
// func (X Strings) Crossover(Y eaopt.Genome, rng *rand.Rand) {
// 	eaopt.CrossGNXString(X, Y.(Strings), 2, rng)
// }

// // MakeStrings creates random Strings slices by picking random characters from a
// // corpus.
// func MakeStrings(rng *rand.Rand) eaopt.Genome {
// 	return Strings(eaopt.InitUnifString(uint(len(dev.Requirements)), corpus, rng))
// }

// // Clone a Strings slice..
// func (X Strings) Clone() eaopt.Genome {
// 	var XX = make(Strings, len(X))
// 	copy(XX, X)
// 	return XX
// }

// func main() {
// 	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	ga.NPops = 4000                  // The number of populations that will be used
// 	ga.PopSize = 3                   // The number of individuals inside each population
// 	ga.NGenerations = 5000           // For many generations the populations will be evolved
// 	ga.HofSize = 2                   // How many of the best individuals should be recorded
// 	ga.Model = eaopt.ModSteadyState{ // Determines how to evolve each population of individuals
// 		Selector:  eaopt.SelElitism{},
// 		MutRate:   0.9,
// 		CrossRate: 0.9,
// 	}
// 	ga.ParallelEval = false

// 	// Add a custom print function to track progress
// 	ga.Callback = func(ga *eaopt.GA) {
// 		// Concatenate the elements from the best individual and display the result
// 		var buffer bytes.Buffer
// 		for _, letter := range ga.HallOfFame[0].Genome.(Strings) {
// 			buffer.WriteString(letter)
// 		}
// 		fmt.Printf("%d) Result -> %s (%.0f mismatches)\n", ga.Generations, buffer.String(), ga.HallOfFame[0].Fitness)
// 	}

// 	// Run the GA
// 	ga.Minimize(MakeStrings)
// }
