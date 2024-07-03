package helpers

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"

	"github.com/google/go-cmp/cmp"
)

// func Map(grid []pkg.Cell, f func(pkg.Cell) pkg.Cell) (res []pkg.Cell) {
func TestMap(t *testing.T) {
	type args struct {
		grid []pkg.Cell
		f    func(pkg.Cell) pkg.Cell
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
				f: func(c pkg.Cell) pkg.Cell {
					c.X = c.X + 1
					c.Y = c.Y + 1
					return c
				},
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(true, 2, 1),
					pkg.NewCell(true, 2, 2),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := Map(tc.a.grid, tc.a.f)
			if diff := cmp.Diff(tc.w.res, got); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
		})
	}
}
