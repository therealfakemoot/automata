package automata

import (
	"log"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

type Neighborhood [9]int

type Grid struct {
	cells         [][]Cell
	left, right   [][]Cell
	Width, Height int
}

func (g *Grid) Wrap(x, y int) (int, int) {
	modularX := mod(x, g.Width)
	log.Printf("modularX(%d) = %d mod %d", modularX, x, g.Width)
	modularY := mod(y, g.Height)
	log.Printf("modularY(%d) = %d mod %d", modularY, y, g.Height)
	return modularX, modularY
}

func (g *Grid) Get(x, y int) Cell {
	log.Printf("getting cell for (%d, %d)\n", x, y)
	modularX, modularY := g.Wrap(x, y)
	return g.cells[modularY][modularX]
}

func (g *Grid) Set(x, y int, c Cell) {
	g.cells[y][x] = c
}

func (g *Grid) Cells() []Cell {
	var cells []Cell
	cells = make([]Cell, 0)
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Width; y++ {
			cells = append(cells, g.cells[y][x])
		}
	}
	return cells
}

func (g *Grid) Neighbors(x, y int) [9]int {
	var neighbors [9]int
	log.Printf("getting neighbors for (%d,%d)\n", x, y)

	neighbors[0] = g.Get(x-1, y+1).State
	neighbors[1] = g.Get(x, y+1).State
	neighbors[2] = g.Get(x+1, y+1).State
	neighbors[3] = g.Get(x-1, y).State
	neighbors[4] = g.Get(x, y).State
	neighbors[5] = g.Get(x+1, y).State
	neighbors[6] = g.Get(x-1, y-1).State
	neighbors[7] = g.Get(x, y-1).State
	neighbors[8] = g.Get(x+1, y-1).State

	return neighbors
}

func NewGrid(width, height int, seed GridSeeder) Grid {
	var g Grid
	g.Width = width
	g.Height = height

	cells := make([][]Cell, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]Cell, width)
	}
	return seed(g)
}
