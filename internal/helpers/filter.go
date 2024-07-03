package helpers

import "chaitanyakolluru/gameOfLife/pkg"

func Filter(grid []pkg.Cell, f func(pkg.Cell) bool) (res []pkg.Cell) {
	for _, c := range grid {
		if f(c) {
			res = append(res, c)
		}
	}
	return
}
