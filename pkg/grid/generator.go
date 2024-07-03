package grid

import (
	"chaitanyakolluru/gameOfLife/pkg"
)

func GridGenerator(x int) ([]pkg.Cell, int) {
	res := []pkg.Cell{}
	for i := 0; i <= x-1; i++ {
		for j := 0; j <= x-1; j++ {
			if (i+j)%2 == 0 {
				res = append(res, pkg.NewCell(true, i, j))
			} else {
				res = append(res, pkg.NewCell(false, i, j))
			}
		}
	}

	return res, x
}
