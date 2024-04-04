package main

import (
	"flag"
	"time"
	// "fmt"

	"github.com/therealfakemoot/automata"
)

func main() {
	var width, height int
	var sleep time.Duration

	flag.IntVar(&width, "width", 100, "horizontal width of automata grid")
	flag.IntVar(&height, "height", 1, "height of automata grid; defaults to 1 for one-dimensional automata")
	flag.DurationVar(&sleep, "cycle", 1*time.Second, "real-time duration of one evaluation cycle")

	flag.Parse()

	trs := automata.TerminalRenderSystem{
		Width: width,
	}
	g := automata.NewGrid(width, height, automata.SeedCenter)
	trs.Grid = &g

	ruleSystem := automata.RuleEngineSystem{
		Grid: trs.Grid,
		Rule: &automata.Rule30{trs.Grid},
	}

	t := time.Tick(sleep)
	for range t {
		ruleSystem.Update(1)
	}
}
