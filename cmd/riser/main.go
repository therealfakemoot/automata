package main

import (
	"math/rand"
	"time"
	// "fmt"

	"github.com/therealfakemoot/automata"
)

func main() {
	trs := automata.TerminalRenderSystem{
		Width: 140,
	}
	trs.Grid = make([]automata.Cell, trs.Width)

	for idx := range trs.Grid {
		trs.Grid[idx].X = idx
		if rand.Float64() > .70 {
			trs.Grid[idx].State = 1
		}
	}

	trs.Grid[3].State = 1

	for {
		trs.Update(1)
		time.Sleep(500 * time.Millisecond)
	}
}
