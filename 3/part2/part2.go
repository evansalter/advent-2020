package part2

import (
	"fmt"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

type Slope struct {
	X int
	Y int
}

var slopes = []Slope{
	{X: 1, Y: 1},
	{X: 3, Y: 1},
	{X: 5, Y: 1},
	{X: 7, Y: 1},
	{X: 1, Y: 2},
}

func Run() {
	lines := helpers.ReadInputFile(3)
	grid := make([][]string, len(lines)-1)
	for i, l := range lines {
		if l == "" {
			continue
		}
		grid[i] = strings.Split(l, "")
	}

	width := len(grid[0])
	finalTally := 1
	for _, s := range slopes {
		numTrees := 0
		i := 0
		for j := 0; j < len(grid); j += s.Y {
			if grid[j][i] == "#" {
				numTrees++
			}
			i = (i + s.X) % width
		}
		finalTally *= numTrees
	}

	fmt.Println(finalTally)
}
