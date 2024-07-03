package rules

import (
	"chaitanyakolluru/gameOfLife/internal/helpers"
	"chaitanyakolluru/gameOfLife/pkg"
)

func ComesBackAlive(grid []pkg.Cell, c pkg.Cell) bool {
	if c.Alive {
		return true
	}

	return len(helpers.Filter(helpers.Neighbors(grid, c),
		func(neh pkg.Cell) bool {
			return neh.Alive
		})) == 3
}
