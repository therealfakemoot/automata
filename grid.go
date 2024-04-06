package automata

import (
	"fmt"
	"strings"
)

type Cell int

func mod(a, b int) int {
	return (a%b + b) % b
}

type Neighborhood [9]int

func (n Neighborhood) String() string {
	var b strings.Builder
	for idx, cell := range n {
		if idx%3 == 0 {
			b.WriteString("\n")
			fmt.Fprint(&b, "|")
		}
		fmt.Fprintf(&b, "%d|", cell)
	}

	return b.String()
}

type Grid struct {
	left, right   [][]Cell
	Width, Height int
	Generation    int
}

func (g *Grid) String() string {
	b := &strings.Builder{}
	for _, row := range g.left {
		for _, c := range row {
			switch {
			case c > 0:
				fmt.Fprint(b, "█")
			case c <= 0:
				fmt.Fprint(b, " ")
			}
		}
	}
	fmt.Fprintf(b, "\n")

	return b.String()
}

func (g *Grid) Wrap(x, y int) (int, int) {
	modularX := mod(x, g.Width)
	modularY := mod(y, g.Height)
	return modularX, modularY
}

func (g *Grid) Get(x, y int) int {
	modularX, modularY := g.Wrap(x, y)
	return int(g.left[modularY][modularX])
}

func (g *Grid) Set(x, y int, c Cell) {
	g.right[y][x] = c
}

func (g *Grid) Cells() []Cell {
	var cells []Cell
	for _, row := range g.left {
		for _, cell := range row {
			cells = append(cells, cell)
		}
	}
	return cells
}

func (g *Grid) Neighbors(x, y int) Neighborhood {
	var neighbors [9]int

	neighbors[0] = g.Get(x-1, y+1)
	neighbors[1] = g.Get(x, y+1)
	neighbors[2] = g.Get(x+1, y+1)

	neighbors[3] = g.Get(x-1, y)
	neighbors[4] = g.Get(x, y)
	neighbors[5] = g.Get(x+1, y)

	neighbors[6] = g.Get(x-1, y-1)
	neighbors[7] = g.Get(x, y-1)
	neighbors[8] = g.Get(x+1, y-1)

	return neighbors
}

func (g *Grid) Update(rule Rule) {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			newCell := rule(x, y, g)
			g.Set(x, y, newCell)
		}
	}
	copy(g.left, g.right)
	g.Generation++
}

func newGrid(w, h int) [][]Cell {
	cells := make([][]Cell, h)
	for y := 0; y < h; y++ {
		cells[y] = make([]Cell, w)
	}

	return cells
}

func NewGrid(width, height int, seed GridSeeder) Grid {
	var g Grid
	g.Width = width
	g.Height = height

	g.left = newGrid(width, height)
	g.right = newGrid(width, height)

	return seed(g)
}
