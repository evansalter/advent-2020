package part2

import (
	"fmt"
	"strconv"

	"github.com/evansalter/advent-2020/helpers"
)

func Run() {
	input := helpers.ReadInputFile(1)
	numbers := make([]int, len(input)-1)
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
		for j, y := range numbers[i:] {
			for _, z := range numbers[j:] {
				if x+y+z == 2020 {
					fmt.Println(x * y * z)
					return
				}
			}
		}
	}
}
