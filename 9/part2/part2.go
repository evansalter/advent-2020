package part2

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/evansalter/advent-2020/helpers"
)

const AnswerFromPart1 = 50047984

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

func CalculateAnswer(nums []int) int {
	sort.Ints(nums)
	return nums[0] + nums[len(nums)-1]
}

func Run() {
	lines := helpers.ReadInputFile(9)
	nums := ParseNums(lines[:len(lines)-1])

	var foundNums []int
	var sum int
	for i := range nums {
		foundNums, sum = make([]int, 0), 0
		for j := i; j < len(lines); j++ {
			num := nums[j]
			foundNums = append(foundNums, num)
			sum += num
			if sum == AnswerFromPart1 {
				ans := CalculateAnswer(foundNums)
				fmt.Println(ans)
				return
			} else if sum > AnswerFromPart1 {
				break
			}
		}
	}

	panic("no solution found")
}
