package main

import (
	"fmt"
	"time"

	"chaitanyakolluru/gameOfLife/pkg"
	"chaitanyakolluru/gameOfLife/pkg/grid"
)

func main() {
	gr, n := grid.GridGenerator(3)
	fmt.Println("Initial grid")
	prettyPrint(gr)
	fmt.Println("Initial grid")

	awesomeThingsHappen(gr, n)
}

func prettyPrint(grid []pkg.Cell) {
	for _, v := range grid {
		fmt.Println(v)
	}
}

func awesomeThingsHappen(g []pkg.Cell, n int) {
	gr2 := grid.ProcessGrid(g, n)
	fmt.Println("+++++++++++++++++++++++")
	prettyPrint(gr2)
	fmt.Println("+++++++++++++++++++++++")

	time.Sleep(1 * time.Second)
	awesomeThingsHappen(gr2, n)
}
