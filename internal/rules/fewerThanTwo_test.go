package rules

import (
	"testing"

	"chaitanyakolluru/gameOfLife/pkg"
)

func TestFewerThanTwo(t *testing.T) {
	type args struct {
		grid []pkg.Cell
		c    pkg.Cell
	}
	type want struct {
		res bool
	}

	type testCondition struct {
		a args
		w want
	}

	cases := map[string]testCondition{
		"should say living when cell is (0,0) and it is a 2x2 grid and all are alive": {
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
				res: true,
			},
		},
		"should say died when cell is (0,0) and it is a 3x3 grid and all except it are dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(false, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 0, 0),
			},
			w: want{
				res: false,
			},
		},
		"should say living when cell is (0,0) and it is a 3x3 grid and (0,1) and (1,0) are alive and all the rest are dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 0, 0),
			},
			w: want{
				res: true,
			},
		},
		"should say died for aliveness when cell is (0,0) and it is a 3x3 grid and (0,1) is alive and all the rest are dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(true, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(false, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 0, 0),
			},
			w: want{
				res: false,
			},
		},
		"should say living when cell is (1,1) and cells (0,1)(1,0)(2,1)(1,2) are alive and the rest dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(true, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(true, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 1, 1),
			},
			w: want{
				res: true,
			},
		},
		"should say living when cell is (1,1) and cells (0,1)(1,0)(2,1) are alive and the rest dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(true, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 1, 1),
			},
			w: want{
				res: true,
			},
		},
		"should say living when cell is (1,1) and cells (0,1)(1,0) are alive and the rest dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(true, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 1, 1),
			},
			w: want{
				res: true,
			},
		},
		"should say died for aliveness when cell is (1,1) and cells (0,1) are alive and the rest dead": {
			a: args{
				grid: []pkg.Cell{
					pkg.NewCell(false, 0, 0),
					pkg.NewCell(true, 0, 1),
					pkg.NewCell(false, 0, 2),
					pkg.NewCell(false, 1, 0),
					pkg.NewCell(true, 1, 1),
					pkg.NewCell(false, 1, 2),
					pkg.NewCell(false, 2, 0),
					pkg.NewCell(false, 2, 1),
					pkg.NewCell(false, 2, 2),
				},
				c: pkg.NewCell(true, 1, 1),
			},
			w: want{
				res: false,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := FewerThanTwo(tc.a.grid, tc.a.c)
			if got != tc.w.res {
				t.Errorf("got: %v, want: %v", got, tc.w.res)
			}
		})
	}
}
