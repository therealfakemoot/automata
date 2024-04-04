package automata

import (
	"fmt"

	"github.com/EngoEngine/ecs"
)

type XYComponent struct {
	X, Y int
}

type AutomataComponent struct {
	State int
}

type Cell struct {
	ecs.BasicEntity
	XYComponent
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
}

func (res *RuleEngineSystem) Remove(e ecs.BasicEntity) {}

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

func (tr *TerminalRenderSystem) Remove(e ecs.BasicEntity) {}

func (tr *TerminalRenderSystem) Priority() int { return -1 }
