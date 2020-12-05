package part2

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
	fmt.Println(findSeat(seatIDs))
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

func findSeat(seatIDs []int) int {
	for i, id := range seatIDs {
		if i != 0 && id-seatIDs[i-1] > 1 {
			return id - 1
		}
		if i != len(seatIDs)-1 && seatIDs[i+1]-id > 1 {
			return id + 1
		}
	}
	return 0
}
