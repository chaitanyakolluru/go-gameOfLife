package helpers

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"

	"github.com/google/go-cmp/cmp"
)

// func Neighbors(grid []pkg.Cell, c pkg.Cell) []pkg.Cell {
func TestNeighbors(t *testing.T) {
	type args struct {
		grid []pkg.Cell
		c    pkg.Cell
	}
	type want struct {
		res []pkg.Cell
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"should exist and take in grid and a cell as args": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
				c: pkg.NewCell(true, 0, 0),
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
			},
		},
		"should return (0,1), (1,0),(1,1) when the input cell is (0,0) and grid is 3x3": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(true, 1, 2),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(true, 2, 1),
					pkg.NewCell(true, 2, 2),
				},
				c: pkg.NewCell(true, 0, 0),
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := Neighbors(tc.a.grid, tc.a.c)
			if diff := cmp.Diff(tc.w.res, got); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
		})
	}
}
