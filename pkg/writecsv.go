package pkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Gen struct {
	Duration    int64         `json:"duration"`
	Generations int64         `json:"generations"`
	HallOfFame  []HallOfFame  `json:"hall_of_fame"`
	Populations []Populations `json:"populations"`
}

type HallOfFame struct {
	Fitness int64    `json:"fitness"`
	Genome  []string `json:"genome"`
	ID      string   `json:"id"`
}

type Indis struct {
	Fitness int64    `json:"fitness"`
	Genome  []string `json:"genome"`
	ID      string   `json:"id"`
}

type Populations struct {
	Rng         struct{} `json:"RNG"`
	Age         int64    `json:"age"`
	Generations int64    `json:"generations"`
	ID          string   `json:"id"`
	Indis       []Indis  `json:"indis"`
}

var i = 0

// WriteFileCsv writes the file.csv
func WriteFileCsv(v interface{}) {
	var gen Gen
	gen = DataMarshaller(v)

	GenDataFile(gen)
	HOFFile(gen.HallOfFame)
	PopulationsFile(gen.Populations)

	return
}

// GenDataFile writes the main csv infos data file.
func GenDataFile(gen Gen) {
	file, err := os.OpenFile(
		"generations.csv",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	defer file.Close()

	fw := csv.NewWriter(file)

	// fw.Write([]string{"Duration", "Generation"})

	var record []string
	record = append(record, fmt.Sprintf("%v", gen.Duration))
	record = append(record, fmt.Sprintf("%v", gen.Generations))
	fw.Write(record)
	fw.Flush()

	if err != nil {
		fmt.Println(err)
	}

	return
}

// HOFFile writes the main csv infos data file.
func HOFFile(hofs []HallOfFame) {
	file, err := os.OpenFile(
		"halloffame.csv",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	defer file.Close()

	fw := csv.NewWriter(file)

	// fw.Write([]string{"ID", "Fitness", "Genome"})

	for _, hof := range hofs {
		var record []string
		record = append(record, fmt.Sprintf("%v", hof.ID))
		record = append(record, fmt.Sprintf("%v", hof.Fitness))
		record = append(record, fmt.Sprintf("%v", hof.Genome))
		fw.Write(record)
	}
	fw.Flush()

	if err != nil {
		fmt.Println(err)
	}

	return
}

// PopulationsFile writes the main csv infos data file.
func PopulationsFile(pops []Populations) {
	file, err := os.OpenFile(
		"populations.csv",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	defer file.Close()

	fw := csv.NewWriter(file)

	// fw.Write([]string{"ID", "Age", "Generations", "Rng"})

	for _, pop := range pops {
		IndisFile(pop.Indis)
		var record []string
		record = append(record, fmt.Sprintf("%v", pop.ID))
		record = append(record, fmt.Sprintf("%v", pop.Age))
		record = append(record, fmt.Sprintf("%v", pop.Generations))
		record = append(record, fmt.Sprintf("%v", pop.Rng))
		fw.Write(record)
	}
	fw.Flush()

	if err != nil {
		fmt.Println(err)
	}

	return
}

// IndisFile writes the main csv infos data file.
func IndisFile(indis []Indis) {
	file, err := os.OpenFile(
		"indis.csv",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	defer file.Close()

	fw := csv.NewWriter(file)

	// fw.Write([]string{"ID", "Fitness"})

	for _, indi := range indis {
		var record []string
		record = append(record, fmt.Sprintf("%v", indi.ID))
		record = append(record, fmt.Sprintf("%v", indi.Fitness))
		fw.Write(record)
		fw.Flush()
	}

	if err != nil {
		fmt.Println(err)
	}

	return
}

// DataMarshaller handles the arriving data.
func DataMarshaller(v interface{}) Gen {
	var gen Gen

	data, err := json.Marshal(v)
	err = json.Unmarshal(data, &gen)

	if err != nil {
		fmt.Println(err)
	}

	return gen
}
