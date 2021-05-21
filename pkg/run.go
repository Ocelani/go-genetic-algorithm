package pkg

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ocelani/go-genetic-algorithm/eaopt"
	"github.com/guptarohit/asciigraph"
)

var dev = NewDevelopment()

var releases = []uint{
	1, 10, 100, 250, 500, 1000, 1500, 2000,
	2500, 3000, 3500, 4000, 4500, 5000,
}

// Run executes the algorithm.
func Run() (b *bytes.Buffer) {
	t := time.Now()
	c := &eaopt.GAConfig{
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
		return nil
	}

	data := []float64{19}

	ga.Callback = func(ga *eaopt.GA) { // Output data
		var buffer bytes.Buffer
		// WriteFileCsv(ga)

		for _, letter := range ga.HallOfFame[0].Genome.(Release) {
			buffer.WriteString(letter) // Concatenate the elements from the best individual
		}
		b = &buffer
		fmt.Printf("\r%v || BestFitness: %.0f || Generation: %v || Running: %v",
			b, ga.HallOfFame[0].Fitness,
			ga.Generations, time.Since(t).Round(time.Millisecond))

		for _, rel := range releases {
			go func(gen, rel uint) {
				if gen == rel {
					data = append(data, ga.HallOfFame[0].Fitness)
					g := asciigraph.Plot(data, asciigraph.Width(50))
					fmt.Printf("\n\n%v\n", g)
				}
			}(ga.Generations, rel)
		}
	}
	ga.Minimize(dev.MakeRelease) // Run the GA
	return b
}

// Finalize prints a conclusion on the console.
func Finalize(b *bytes.Buffer) {
	fmt.Printf("\n\nSTAKEHOLDERS")

	for _, stk := range dev.Stakeholders {
		var reqs, pris string
		for _, r := range stk.Requirements {
			reqs = reqs + " " + r.Char
			pris = pris + " " + strconv.Itoa(r.Priority)
		}
		fmt.Printf("\n ID: %d \n PRIORITY: %d\n%s\n%s\n",
			stk.ID, stk.Priority, reqs, pris)
	}

	var reqs, pris, stks string
	for _, r := range dev.Requirements {
		stks = stks + " " + strconv.Itoa(r.StakeID)
		reqs = reqs + " " + r.Char
		pris = pris + " " + strconv.Itoa(r.Priority)
	}
	fmt.Printf("\nDEVELOPMENT\nSTK\t%s\nREQ\t%s\nPRI\t%s\n", stks, reqs, pris)

	want := strings.ReplaceAll(reqs, " ", "")
	got := b.String()

	var found, w, g string
	for i, char := range got {
		if !strings.Contains(string(want[i]), string(char)) {
			found = found + " " + "✘"
		} else {
			found = found + " " + "✔"
		}
		w = w + " " + string(want[i])
		g = g + " " + string(char)
	}
	fmt.Printf("\nREQUIRE\t%s\nRELEASE\t%s\nSUCCESS\t%s\n\n", w, g, found)
}
