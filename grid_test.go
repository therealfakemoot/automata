package automata_test

import (
	"testing"

	"github.com/therealfakemoot/automata"
)

func Test_1DTorusCoordinates(t *testing.T) {
	cases := []struct {
		name       string
		inX, inY   int
		outX, outY int
	}{
		{
			name: "top-left",
			inX:  -1,
			inY:  -1,
			outX: 0,
			outY: 0,
		},
	}
	// this test case uses a 1d plane 10 units wide and 1 unit tall (duh)
	g := automata.NewGrid(10, 1, automata.SeedConstant(0))
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actualX, actualY := g.Wrap(tc.inX, tc.inY)
			if actualX != tc.outX {
				t.Log("actualX did not match")
				t.Fail()
			}
			if actualY != tc.outY {
				t.Log("actualY did not match")
				t.Fail()
			}
		})
	}
}