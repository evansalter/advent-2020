package part1

import (
	"fmt"
	"strconv"

	"github.com/evansalter/advent-2020/helpers"
)

const PreambleLength = 25

func ParseNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return num
}

func ParseNums(strs []string) []int {
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i] = ParseNum(s)
	}
	return nums
}

func SumExists(nums []int, target int) bool {
	for i, numA := range nums {
		for j := i + 1; j < len(nums); j++ {
			numB := nums[j]
			if numA == numB {
				continue
			}
			if numA+numB == target {
				return true
			}
		}
	}
	return false
}

func Run() {
	lines := helpers.ReadInputFile(9)
	for i := PreambleLength; i < len(lines)-1; i++ {
		cur := ParseNum(lines[i])
		prev := ParseNums(lines[i-PreambleLength : i])
		if !SumExists(prev, cur) {
			fmt.Println(cur)
			return
		}
	}

	panic("no solution found")
}
