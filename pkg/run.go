package pkg

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
	"github.com/guptarohit/asciigraph"
)

// Run executes the algorithm.
func Run() {
	t := time.Now()
	c := &eaopt.GAConfig{
		RNG:          rand.New(rand.NewSource(time.Now().UnixNano())),
		NPops:        400,  // The number of populations that will be used
		PopSize:      3,    // The number of individuals inside each population
		NGenerations: 5001, // For many generations the populations will be evolved
		HofSize:      300,  // How many of the best individuals should be recorded
		Model: eaopt.ModSteadyState{
			Selector:  eaopt.SelElitism{},
			MutRate:   0.1,
			CrossRate: 0.9,
			KeepBest:  true,
		},
	}
	ga, err := c.NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []float64{19}

	ga.Callback = func(ga *eaopt.GA) { // Output data
		var buffer bytes.Buffer
		// WriteFileCsv(ga)

		for _, letter := range ga.HallOfFame[0].Genome.(Release) {
			buffer.WriteString(letter) // Concatenate the elements from the best individual
		}
		fmt.Printf("\r %v - Best fitness at generation %d: %.0f",
			time.Since(t), ga.Generations, ga.HallOfFame[0].Fitness)

		go func(gen uint) {
			if gen == 1 || gen == 10 || gen == 50 || gen == 100 ||
				gen == 500 || gen == 1000 || gen == 1500 || gen == 2000 || gen == 2500 ||
				gen == 3000 || gen == 3500 || gen == 4000 || gen == 4500 || gen == 5000 {

				data = append(data, ga.HallOfFame[0].Fitness)
				g := asciigraph.Plot(data, asciigraph.Width(50))
				fmt.Println()
				fmt.Println(g)
				fmt.Println()
			}
		}(ga.Generations)
	}
	ga.Minimize(dev.MakeRelease) // Run the GA
}
