package helpers

import (
	"chaitanyakolluru/gameOfLife/pkg"
)

func Neighbors(grid []pkg.Cell, c pkg.Cell) []pkg.Cell {
	return Filter(grid, func(cc pkg.Cell) bool {
		d := DistanceFrom(cc, c)
		return d < 2 && d > 0
	})
}
