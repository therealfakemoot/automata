package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/therealfakemoot/automata"
)

func main() {
	var width, height int
	var sleep time.Duration

	flag.IntVar(&width, "width", 100, "horizontal width of automata grid")
	flag.IntVar(&height, "height", 1, "height of automata grid; defaults to 1 for one-dimensional automata")
	flag.DurationVar(&sleep, "cycle", 1*time.Second, "real-time duration of one evaluation cycle")

	flag.Parse()

	g := automata.NewGrid(width, height, automata.SeedCenter)

	t := time.Tick(sleep)
	for range t {
		fmt.Printf("%s", g.String())
		g.Update(automata.Rule30)
	}
}
