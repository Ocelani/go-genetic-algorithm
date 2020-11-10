package pkg

import (
	"fmt"

	eaopt "github.com/MaxHalford/eaopt"
)

// Run executes the algorithm.
func Run() {
	var ga eaopt.GA

	// Instantiate a GA with a GAConfig
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the number of generations to run for
	ga.NGenerations = 10

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Find the minimum
	err = ga.Minimize(VectorFactory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
