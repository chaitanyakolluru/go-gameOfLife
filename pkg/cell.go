package pkg

type Cell struct {
	Alive bool
	X     int
	Y     int
}

func NewCell(alive bool, x, y int) Cell {
	return Cell{
		Alive: alive,
		X:     x,
		Y:     y,
	}
}
