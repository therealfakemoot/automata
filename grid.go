package automata

import (
	// "log"
	"math"
	"math/rand"
)

type GridSeeder func(Grid) Grid

func SeedIdentity(grid Grid) Grid {
	return grid
}

func SeedConstant(c int) GridSeeder {
	return func(grid Grid) Grid {
		for idx := range grid.Cells {
			grid.Cells[idx].State = c
		}
		return grid
	}
}

func SeedRandom(x float64) GridSeeder {
	return func(grid Grid) Grid {
		for idx := range grid.Cells {
			grid.Cells[idx].X = idx
			if rand.Float64() >= x {
				grid.Cells[idx].State = 1
			}
		}
		return grid
	}
}

type Grid struct {
	Cells []Cell
	Width int
}

type Neighborhood [9]int

func (g *Grid) Neighbors(x, y int) [9]int {
	var neighbors [9]int

	neighbors[0] = g.Cells[((x-1)*g.Width)+(y+1)].State
	neighbors[1] = g.Cells[((x)*g.Width)+(y+1)].State
	neighbors[2] = g.Cells[((x+1)*g.Width)+(y+1)].State
	neighbors[3] = g.Cells[((x-1)*g.Width)+(y)].State
	neighbors[4] = g.Cells[((x)*g.Width)+(y)].State
	neighbors[5] = g.Cells[((x+1)*g.Width)+(y)].State
	neighbors[6] = g.Cells[((x-1)*g.Width)+(y-1)].State
	neighbors[7] = g.Cells[((x)*g.Width)+(y-1)].State
	neighbors[8] = g.Cells[((x+1)*g.Width)+(y-1)].State

	return neighbors
}

func NewGrid(x, y int, seed GridSeeder) Grid {
	var g Grid
	gridLen := int(math.Max(float64(x)*float64(y), float64(x)))
	grid := make([]Cell, gridLen)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			unwrappedIdx := (i * x) + j
			grid[unwrappedIdx].X = i
			grid[unwrappedIdx].Y = i
		}
	}
	g.Cells = grid
	return seed(g)
}

type Rule interface {
	Apply(Cell) Cell
}
