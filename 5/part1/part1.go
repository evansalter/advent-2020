package part1

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

func Run() {
	lines := helpers.ReadInputFile(5)
	seatIDs := make([]int, len(lines)-1)
	for i, l := range lines {
		if l == "" {
			continue
		}
		seatIDs[i] = getSeatID(l)
	}

	sort.Ints(seatIDs)
	fmt.Println(seatIDs[len(seatIDs)-1])
}

func getSeatID(in string) int {
	rowMin, rowMax, colMin, colMax := 0, 127, 0, 7
	chars := strings.Split(in, "")
	for _, c := range chars {
		rowDiff := rowMax - rowMin
		colDiff := colMax - colMin
		switch c {
		case "F":
			rowMax -= int(math.Floor(float64(rowDiff) / 2.0))
		case "B":
			rowMin += int(math.Ceil(float64(rowDiff) / 2.0))
		case "R":
			colMin += int(math.Ceil(float64(colDiff) / 2.0))
		case "L":
			colMax -= int(math.Floor(float64(colDiff) / 2.0))
		}
	}

	return (rowMin * 8) + colMin
}
