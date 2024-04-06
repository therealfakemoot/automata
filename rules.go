package automata

import (
	"log"
)

type Rule func(x, y int, g *Grid) Cell

func Rule30(x, y int, g *Grid) Cell {
	var (
		case1 = Neighborhood{
			1, 1, 1,
			1, 1, 1,
			1, 1, 1,
		}
		case2 = Neighborhood{
			1, 1, 0,
			1, 1, 0,
			1, 1, 0,
		}
		case3 = Neighborhood{
			1, 0, 1,
			1, 0, 1,
			1, 0, 1,
		}
		case4 = Neighborhood{
			1, 0, 0,
			1, 0, 0,
			1, 0, 0,
		}
		case5 = Neighborhood{
			0, 1, 1,
			0, 1, 1,
			0, 1, 1,
		}
		case6 = Neighborhood{
			0, 1, 0,
			0, 1, 0,
			0, 1, 0,
		}
		case7 = Neighborhood{
			0, 0, 1,
			0, 0, 1,
			0, 0, 1,
		}
		case8 = Neighborhood{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		}
	)

	var state int
	neighbors := g.Neighbors(x, y)
	switch {
	case neighbors == case1:
		state = 0
	case neighbors == case2:
		state = 1
	case neighbors == case3:
		state = 0
	case neighbors == case4:
		state = 1
	case neighbors == case5:
		state = 1
	case neighbors == case6:
		state = 1
	case neighbors == case7:
		state = 1
	case neighbors == case8:
		state = 0
	default:
		log.Println("didn't apply any rules")
	}

	return Cell(state)
}
