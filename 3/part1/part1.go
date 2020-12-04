package part1

import (
	"fmt"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

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
	numTrees := 0
	i := 0
	for j := 0; j < len(grid); j++ {
		if grid[j][i] == "#" {
			numTrees++
		}
		i = (i + 3) % width
	}

	fmt.Println(numTrees)
}
