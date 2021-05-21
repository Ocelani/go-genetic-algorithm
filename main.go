package main

import "github.com/Ocelani/go-genetic-algorithm/pkg"

// main just initializes and finalizes the program.
func main() {
	b := pkg.Run()
	pkg.Finalize(b)
}
