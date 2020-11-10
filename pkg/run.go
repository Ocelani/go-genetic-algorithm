package pkg

import (
	"bytes"
	"fmt"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
)

// Run executes the algorithm.
func Run() {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}
	ga.NGenerations = 30

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
}
