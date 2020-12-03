package part1

import (
	"fmt"
	"strconv"

	"github.com/evansalter/advent-2020/helpers"
)

func Run() {
	input := helpers.ReadInputFile(1)
	numbers := make([]int, len(input))
	for i, line := range input {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err.Error())
		}
		numbers[i] = num
	}

	for i, x := range numbers {
		for _, y := range numbers[i:] {
			if x+y == 2020 {
				fmt.Println(x * y)
				return
			}
		}
	}
}
