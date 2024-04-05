package automata

func mod(a, b int) int {
	return (a%b + b) % b
}

type Neighborhood [9]int

type Grid struct {
	left, right   [][]Cell
	Width, Height int
}

func (g *Grid) Wrap(x, y int) (int, int) {
	modularX := mod(x, g.Width)
	modularY := mod(y, g.Height)
	return modularX, modularY
}

func (g *Grid) Get(x, y int) Cell {
	modularX, modularY := g.Wrap(x, y)
	return g.left[modularY][modularX]
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

func (g *Grid) Neighbors(x, y int) [9]int {
	var neighbors [9]int

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
