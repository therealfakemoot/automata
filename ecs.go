package automata

import (
	"fmt"
)

type XYComponent struct {
	X, Y int
}

type AutomataComponent struct {
	State int
}

type Cell struct {
	AutomataComponent
}

type RuleEngineSystem struct {
	Rule Rule
	Grid *Grid
}

func (res *RuleEngineSystem) Update(dt float32) {
	for x := 0; x < res.Grid.Width; x++ {
		for y := 0; y < res.Grid.Height; y++ {
			res.Grid.Set(x, y, res.Grid.Get(x, y))
		}
	}
	res.Grid.left, res.Grid.right = res.Grid.right, res.Grid.left
	for _, c := range res.Grid.Cells() {
		fmt.Printf("%d", c.State)
	}
	fmt.Printf("\n")
}

func (res *RuleEngineSystem) Priority() int { return 0 }

type TerminalRenderSystem struct {
	Grid  *Grid
	Width int
}

func (trs *TerminalRenderSystem) Update(dt float32) {
	for _, c := range trs.Grid.Cells() {
		fmt.Printf("%d", c.State)
	}
	fmt.Printf("\n")
}

func (tr *TerminalRenderSystem) Priority() int { return -1 }
