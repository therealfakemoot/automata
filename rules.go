package automata

type Rule interface {
	Apply(int, int) Cell
}

type Rule30 struct {
	*Grid
}

func (r *Rule30) Apply(x, y int) Cell {
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

	cell := Cell{}
	neighbors := r.Neighbors(x, y)
	switch {
	case neighbors == case1:
		cell.State = 0
	case neighbors == case2:
		cell.State = 0
	case neighbors == case3:
		cell.State = 0
	case neighbors == case4:
		cell.State = 1
	case neighbors == case5:
		cell.State = 1
	case neighbors == case6:
		cell.State = 1
	case neighbors == case7:
		cell.State = 1
	case neighbors == case8:
		cell.State = 0
	}

	return cell
}
