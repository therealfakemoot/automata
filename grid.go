package automata

import (
	// "log"
	"math"
	"math/rand"
)

type Grid struct {
	Width int
	Cells []Cell
	Rules []Rule
}

func (g *Grid) Update() {
	for cellIdx, cell := range g.Cells {
		for _, rule := range g.Rules {
			g.Cells[cellIdx] = rule.Apply(cell)
		}
	}
}

type GridSeeder func([]Cell) []Cell

func SeedIdentity(grid []Cell) []Cell {
	return grid
}

func SeedConstant(c int) GridSeeder {
	return func(grid []Cell) []Cell {
		for idx := range grid {
			grid[idx].State = c
		}
		return grid
	}
}

func SeedRandom(x float64) GridSeeder {
	return func(grid []Cell) []Cell {
		for idx := range grid {
			grid[idx].X = idx
			if rand.Float64() >= x {
				grid[idx].State = 1
			}
		}
		return grid
	}
}

func NewGrid(x, y int, seed GridSeeder) []Cell {
	gridLen := int(math.Max(float64(x)*float64(y), float64(x)))
	grid := make([]Cell, gridLen)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			unwrappedIdx := (i * x) + j
			grid[unwrappedIdx].X = i
			grid[unwrappedIdx].Y = i
		}
	}

	return seed(grid)
}

type Rule interface {
	Apply(Cell) Cell
}
