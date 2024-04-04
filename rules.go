package automata

type Rule30 struct {
	Grid
}

func (r *Rule30) Apply(source Cell) Cell {
	/*
		left, right := neighbors[3], neighors[7]

		switch source.State {
		case 0:
		}
	*/

	case1 := [9]int{
		0, 0, 0,
		1, 1, 1,
		0, 0, 0,
	}

	neighbors := r.Neighbors(source.X, source.Y)
	switch {
	case neighbors == case1:
		source.State = 0
	}

	return source
}
