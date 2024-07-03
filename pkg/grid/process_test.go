package grid

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"

	"github.com/google/go-cmp/cmp"
)

// func ProcessGrid(grid []pkg.Cell, x int) (res []pkg.Cell) {
func TestProcessGrid(t *testing.T) {
	type args struct {
		grid []pkg.Cell
		x    int
	}
	type want struct {
		res []pkg.Cell
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"should take in a grid and return a grid only with alive elements": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
				x: 2,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
			},
		},
		"sjklj;hould take in a grid and return a grid only with alive elements": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
				x: 2,
			},
			w: want{
				res: nil,
			},
		},
		"should process a 3x3 grid and return a grid with alive cells; ex1": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(true, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				x: 3,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(true, 1, 2),
					pkg.NewCell(true, 2, 1),
				},
			},
		},
		"should process a 4x4 grid and return a grid with alive cells; ex2": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(false, 0, 3),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(true, 1, 2),
					pkg.NewCell(true, 1, 3),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
					pkg.NewCell(true, 2, 3),
					pkg.NewCell(true, 3, 0),
					pkg.NewCell(true, 3, 1),
					pkg.NewCell(false, 3, 2),
					pkg.NewCell(true, 3, 3),
				},
				x: 4,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(true, 1, 2),
					pkg.NewCell(true, 1, 3),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(true, 2, 3),
					pkg.NewCell(true, 3, 2),
				},
			},
		},
		"should process a 5x5 grid and return a grid with alive cells; ex3": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(false, 0, 3),
					pkg.NewCell(true, 0, 4),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(true, 1, 3),
					pkg.NewCell(false, 1, 4),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(true, 2, 2),
					pkg.NewCell(false, 2, 3),
					pkg.NewCell(true, 2, 4),
					pkg.NewCell(false, 3, 0),
					pkg.NewCell(true, 3, 1),
					pkg.NewCell(false, 3, 2),
					pkg.NewCell(true, 3, 3),
					pkg.NewCell(false, 3, 4),
					pkg.NewCell(true, 4, 0),
					pkg.NewCell(false, 4, 1),
					pkg.NewCell(true, 4, 2),
					pkg.NewCell(false, 4, 3),
					pkg.NewCell(true, 4, 4),
				},
				x: 5,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 0, 3),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 4),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(true, 2, 4),
					pkg.NewCell(true, 3, 0),
					pkg.NewCell(true, 3, 4),
					pkg.NewCell(true, 4, 1),
					pkg.NewCell(true, 4, 2),
					pkg.NewCell(true, 4, 3),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := ProcessGrid(tc.a.grid, tc.a.x)
			if diff := cmp.Diff(tc.w.res, got); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
		})
	}
}
