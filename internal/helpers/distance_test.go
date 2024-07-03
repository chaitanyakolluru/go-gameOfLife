package helpers

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"
)

// func DistanceFrom(a, b pkg.Cell) float64 {
func TestDistanceFrom(t *testing.T) {
	type args struct {
		a pkg.Cell
		b pkg.Cell
	}
	type want struct {
		res int64
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"should give distance between two points (0,0) and (0,1) as 1": {
			a: args{
				a: pkg.NewCell(true, 0, 0),
				b: pkg.NewCell(false, 0, 1),
			},
			w: want{
				res: 1,
			},
		},
		"should give distance between two points (0,0) and (0,3) as 3": {
			a: args{
				a: pkg.NewCell(true, 0, 0),
				b: pkg.NewCell(false, 0, 3),
			},
			w: want{
				res: 3,
			},
		},
		"should give distance between two points (0,2) and (2,0) as 2": {
			a: args{
				a: pkg.NewCell(true, 0, 2),
				b: pkg.NewCell(false, 2, 0),
			},
			w: want{
				res: 2,
			},
		},
		"should give distance between two points (1,5) and (4,5) as 3": {
			a: args{
				a: pkg.NewCell(true, 1, 5),
				b: pkg.NewCell(false, 4, 5),
			},
			w: want{
				res: 3,
			},
		},
		"should give distance between two points (4,5) and (1,5) as 3": {
			a: args{
				a: pkg.NewCell(true, 4, 5),
				b: pkg.NewCell(false, 1, 5),
			},
			w: want{
				res: 3,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := DistanceFrom(tc.a.a, tc.a.b)
			if got != tc.w.res {
				t.Errorf("got: %v, want: %v", got, tc.w.res)
			}
		})
	}
}
