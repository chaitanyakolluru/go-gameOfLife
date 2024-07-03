package rules

import (
	"chaitanyakolluru/gameOfLife/internal/helpers"
	"chaitanyakolluru/gameOfLife/pkg"
)

func TwoOrThree(grid []pkg.Cell, c pkg.Cell) bool {
	aliveCells := len(helpers.Filter(helpers.Neighbors(grid, c),
		func(neh pkg.Cell) bool {
			return neh.Alive
		}))

	return aliveCells == 2 || aliveCells == 3
}
