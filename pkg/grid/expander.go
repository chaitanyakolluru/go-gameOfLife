package grid

import (
	"chaitanyakolluru/gameOfLife/internal/helpers"
	"chaitanyakolluru/gameOfLife/pkg"
)

func GridExpander(grid []pkg.Cell, gSize int) (res []pkg.Cell) {
	var i, j int
	for i = 0; i <= gSize-1; {
		for j = 0; j <= gSize-1; {
			tempRes := helpers.Filter(grid, func(c pkg.Cell) bool {
				return c.X == i && c.Y == j
			})

			if len(tempRes) == 0 {
				res = append(res, pkg.NewCell(false, i, j))
			} else {
				res = append(res, tempRes...)
			}
			j++
		}
		i++
		j = 0
	}

	return
}
