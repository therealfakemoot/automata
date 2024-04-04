package automata

import (
	"math/rand"
)

type GridSeeder func(Grid) Grid

func SeedIdentity(grid Grid) Grid {
	return grid
}

func SeedCenter(grid Grid) Grid {
	grid.cells[len(grid.cells)/2][len(grid.cells)/2].State = 1
	return grid
}

func SeedConstant(c int) GridSeeder {
	return func(grid Grid) Grid {
		for y, row := range grid.cells {
			for x := range row {
				grid.cells[y][x].State = c
			}
		}
		return grid
	}
}

func SeedRandom(seed float64) GridSeeder {
	return func(grid Grid) Grid {
		for y, row := range grid.cells {
			for x := range row {
				if rand.Float64() >= seed {
					grid.cells[y][x].State = 1
				}
			}
		}
		return grid
	}
}
