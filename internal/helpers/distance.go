package helpers

import (
	"math"

	"chaitanyakolluru/gameOfLife/pkg"
)

func DistanceFrom(a, b pkg.Cell) int64 {
	return int64(math.Abs(math.Hypot(float64(b.X)-float64(a.X), float64(b.Y)-float64(a.Y))))
}
