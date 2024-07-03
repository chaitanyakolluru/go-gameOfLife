package rules

import (
	"chaitanyakolluru/gameOfLife/internal/helpers"
	"chaitanyakolluru/gameOfLife/pkg"
)

func FewerThanTwo(grid []pkg.Cell, c pkg.Cell) bool {
	return !(len(helpers.Filter(helpers.Neighbors(grid, c),
		func(neh pkg.Cell) bool {
			return neh.Alive
		})) < 2)
}
