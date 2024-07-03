package grid

import (
	"chaitanyakolluru/gameOfLife/internal/helpers"
	"chaitanyakolluru/gameOfLife/internal/rules"
	"chaitanyakolluru/gameOfLife/pkg"
)

func ProcessGrid(grid []pkg.Cell, x int) []pkg.Cell {
	res := helpers.Filter(GridExpander(grid, x), func(c pkg.Cell) bool {
		return rules.FewerThanTwo(grid, c) &&
			rules.TwoOrThree(grid, c) &&
			rules.MoreThanThree(grid, c) &&
			rules.ComesBackAlive(grid, c)
	})

	return helpers.Map(res, func(c pkg.Cell) pkg.Cell {
		c.Alive = true
		return c
	})
}
