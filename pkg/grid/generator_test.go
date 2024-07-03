package grid

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"

	"github.com/google/go-cmp/cmp"
)

// func GridGenerator(x int) ([]pkg.Cell, int) {
func TestGenerator(t *testing.T) {
	type args struct {
		x int
	}
	type want struct {
		res []pkg.Cell
		y   int
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"should take in an input as a seed and generate a grid with cells whose x and y coordinates sum upto an even number": {
			a: args{
				x: 2,
			},
			w: want{
				res: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(true, 1, 1),
				},
				y: 2,
			},
		},
		"with 5 should take in an input as a seed and generate a grid with cells whose x and y coordinates sum upto an even number": {
			a: args{
				x: 5,
			},
			w: want{
				res: []pkg.Cell{
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
				y: 5,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got1, got2 := GridGenerator(tc.a.x)
			if diff := cmp.Diff(tc.w.res, got1); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
			if diff := cmp.Diff(tc.w.y, got2); diff != "" {
				t.Errorf("\n%v\n -want, +got:\n%v\n", tc.w.res, diff)
			}
		})
	}
}
