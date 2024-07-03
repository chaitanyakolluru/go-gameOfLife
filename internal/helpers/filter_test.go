package helpers

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"

	"github.com/google/go-cmp/cmp"
)

// func Filter(grid []pkg.Cell, f func(pkg.Cell) bool) (res []pkg.Cell) {
func TestFilter(t *testing.T) {
	type args struct {
		grid []pkg.Cell
		f    func(pkg.Cell) bool
	}
	type want struct {
		res []pkg.Cell
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"test 1": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
				f: func(c pkg.Cell) bool {
					return !c.Alive
				},
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(false, 0, 1),
				},
			},
		},
		"test 2": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				f: func(c pkg.Cell) bool {
					return !c.Alive
				},
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := Filter(tc.a.grid, tc.a.f)
			if diff := cmp.Diff(tc.w.res, got); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
		})
	}
}
