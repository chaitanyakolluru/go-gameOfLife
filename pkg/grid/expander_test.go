package grid

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"

	"github.com/google/go-cmp/cmp"
)

// func GridExpander(grid []pkg.Cell, gSize int) (res []pkg.Cell) {
func TestExpander(t *testing.T) {
	type args struct {
		grid []pkg.Cell
		size int
	}
	type want struct {
		res []pkg.Cell
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"should expand a grid if some of the cells are missing from it - size 2": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
				},
				size: 2,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(false, 1, 1),
				},
			},
		},
		"should expand a grid and add dead cells if some are missing from it - size 5": {
			a: args{
				grid: []pkg.Cell{
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
				size: 5,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 0, 3),
					pkg.NewCell(false, 0, 4),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 1, 3),
					pkg.NewCell(true, 1, 4),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
					pkg.NewCell(false, 2, 3),
					pkg.NewCell(true, 2, 4),
					pkg.NewCell(true, 3, 0),
					pkg.NewCell(false, 3, 1),
					pkg.NewCell(false, 3, 2),
					pkg.NewCell(false, 3, 3),
					pkg.NewCell(true, 3, 4),
					pkg.NewCell(false, 4, 0),
					pkg.NewCell(true, 4, 1),
					pkg.NewCell(true, 4, 2),
					pkg.NewCell(true, 4, 3),
					pkg.NewCell(false, 4, 4),
				},
			},
		},
		"should'nt change in any way a grid that's already expanded - size 5": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 0, 3),
					pkg.NewCell(false, 0, 4),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 1, 3),
					pkg.NewCell(true, 1, 4),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
					pkg.NewCell(false, 2, 3),
					pkg.NewCell(true, 2, 4),
					pkg.NewCell(true, 3, 0),
					pkg.NewCell(false, 3, 1),
					pkg.NewCell(false, 3, 2),
					pkg.NewCell(false, 3, 3),
					pkg.NewCell(true, 3, 4),
					pkg.NewCell(false, 4, 0),
					pkg.NewCell(true, 4, 1),
					pkg.NewCell(true, 4, 2),
					pkg.NewCell(true, 4, 3),
					pkg.NewCell(false, 4, 4),
				},
				size: 5,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(true, 0, 2),
					pkg.NewCell(true, 0, 3),
					pkg.NewCell(false, 0, 4),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 1, 3),
					pkg.NewCell(true, 1, 4),
					pkg.NewCell(true, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
					pkg.NewCell(false, 2, 3),
					pkg.NewCell(true, 2, 4),
					pkg.NewCell(true, 3, 0),
					pkg.NewCell(false, 3, 1),
					pkg.NewCell(false, 3, 2),
					pkg.NewCell(false, 3, 3),
					pkg.NewCell(true, 3, 4),
					pkg.NewCell(false, 4, 0),
					pkg.NewCell(true, 4, 1),
					pkg.NewCell(true, 4, 2),
					pkg.NewCell(true, 4, 3),
					pkg.NewCell(false, 4, 4),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := GridExpander(tc.a.grid, tc.a.size)
			if diff := cmp.Diff(tc.w.res, got); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
		})
	}
}
