package automata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
			inY:  +1,
			outX: 9,
			outY: 0,
		},
		{
			name: "top",
			inX:  0,
			inY:  1,
			outX: 0,
			outY: 0,
		},
		{
			name: "top-right",
			inX:  1,
			inY:  1,
			outX: 1,
			outY: 0,
		},
		{
			name: "left",
			inX:  -1,
			inY:  0,
			outX: 9,
			outY: 0,
		},
		{
			name: "identity",
			inX:  0,
			inY:  0,
			outX: 0,
			outY: 0,
		},
		{
			name: "right",
			inX:  1,
			inY:  0,
			outX: 1,
			outY: 0,
		},
		{
			name: "bottom-left",
			inX:  -1,
			inY:  -1,
			outX: 9,
			outY: 0,
		},
		{
			name: "bottom",
			inX:  0,
			inY:  -1,
			outX: 0,
			outY: 0,
		},
		{
			name: "bottom-right",
			inX:  1,
			inY:  -1,
			outX: 1,
			outY: 0,
		},
	}

	// this test case uses a 1d plane 10 units wide and 1 unit tall (duh)
	// test cases are centered on (0,0)
	g := automata.NewGrid(10, 1, automata.SeedConstant(0))
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actualX, actualY := g.Wrap(tc.inX, tc.inY)
			if actualX != tc.outX {
				t.Logf("actualX did not match: expected %d, got %d", tc.outX, actualX)
				t.Fail()
			}
			if actualY != tc.outY {
				t.Logf("actualY did not match: expected %d, got %d", tc.outY, actualY)
				t.Fail()
			}
		})
	}
}

func Test_GridNeighors(t *testing.T) {
	grid := automata.NewGrid(
		10, 1,
		automata.SeedCenter,
	)
	cases := []struct {
		name     string
		x, y     int
		expected [9]int
	}{
		{
			name: "origin",
			x:    0,
			y:    0,
			expected: [9]int{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			},
		},
		{
			name: "seed",
			x:    5,
			y:    0,
			expected: [9]int{
				0, 1, 0,
				0, 1, 0,
				0, 1, 0,
			},
		},
		{
			name: "left-of-seed",
			x:    4,
			y:    0,
			expected: [9]int{
				0, 0, 1,
				0, 0, 1,
				0, 0, 1,
			},
		},
		{
			name: "right-of-seed",
			x:    6,
			y:    0,
			expected: [9]int{
				1, 0, 0,
				1, 0, 0,
				1, 0, 0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			neighbors := grid.Neighbors(tc.x, tc.y)
			assert.EqualValues(t, neighbors, tc.expected)
		})
	}
}
