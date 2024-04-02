package main

import (
	"flag"
	"time"
	// "fmt"

	"github.com/EngoEngine/ecs"
	"github.com/therealfakemoot/automata"
)

func main() {
	var width, height int
	var sleep time.Duration

	flag.IntVar(&width, "width", 100, "horizontal width of automata grid")
	flag.IntVar(&height, "height", 0, "height of automata grid; defaults to 0 for one-dimensional automata")
	flag.DurationVar(&sleep, "cycle", 1*time.Second, "real-time duration of one evaluation cycle")

	world := ecs.World{}
	trs := automata.TerminalRenderSystem{
		Width: width,
	}
	trs.Grid = automata.NewGrid(width, height, automata.SeedConstant(1))
	world.AddSystem(&trs)

	t := time.Tick(sleep)
	for range t {
		world.Update(1)
	}
}
