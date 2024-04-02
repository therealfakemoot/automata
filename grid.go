package automata

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

type Rule interface {
	Apply(Cell) Cell
}
