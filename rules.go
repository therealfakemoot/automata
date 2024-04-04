package automata

import (
	"log"
)

type Rule30 struct {
	*Grid
}

func (r *Rule30) Apply(source Cell) Cell {
	log.Printf("%#v\n", r)

	var (
		case1 = [9]int{
			0, 0, 0,
			1, 1, 1,
			0, 0, 0,
		}
		case2 = [9]int{
			0, 0, 0,
			1, 1, 0,
			0, 0, 0,
		}
		case3 = [9]int{
			0, 0, 0,
			1, 0, 1,
			0, 0, 0,
		}
		case4 = [9]int{
			0, 0, 0,
			1, 0, 0,
			0, 0, 0,
		}
		case5 = [9]int{
			0, 0, 0,
			0, 1, 1,
			0, 0, 0,
		}
		case6 = [9]int{
			0, 0, 0,
			0, 1, 0,
			0, 0, 0,
		}
		case7 = [9]int{
			0, 0, 0,
			0, 0, 1,
			0, 0, 0,
		}
		case8 = [9]int{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		}
	)

	neighbors := r.Neighbors(source.X, source.Y)
	switch {
	case neighbors == case1:
		source.State = 0
	case neighbors == case2:
		source.State = 0
	case neighbors == case3:
		source.State = 0
	case neighbors == case4:
		source.State = 1
	case neighbors == case5:
		source.State = 1
	case neighbors == case6:
		source.State = 1
	case neighbors == case7:
		source.State = 1
	case neighbors == case8:
		source.State = 0
	}

	return source
}
