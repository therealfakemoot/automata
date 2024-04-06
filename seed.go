package automata

import (
	"math/rand"
)

type GridSeeder func(Grid) Grid

func SeedIdentity(grid Grid) Grid {
	return grid
}

func SeedCenter(grid Grid) Grid {
	grid.right[0][grid.Width/2] = 1
	grid.left, grid.right = grid.right, grid.left
	return grid
}

func SeedConstant(c int) GridSeeder {
	return func(grid Grid) Grid {
		for y, row := range grid.left {
			for x := range row {
				grid.right[y][x] = Cell(c)
			}
		}
		grid.left, grid.right = grid.right, grid.left
		return grid
	}
}

func SeedRandom(seed float64) GridSeeder {
	return func(grid Grid) Grid {
		for y, row := range grid.left {
			for x := range row {
				if rand.Float64() >= seed {
					grid.right[y][x] = 1
				}
			}
		}
		grid.left, grid.right = grid.right, grid.left
		return grid
	}
}
