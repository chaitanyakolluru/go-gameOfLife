package helpers

import "chaitanyakolluru/gameOfLife/pkg"

func Map(grid []pkg.Cell, f func(pkg.Cell) pkg.Cell) (res []pkg.Cell) {
	for _, c := range grid {
		res = append(res, f(c))
	}
	return
}
