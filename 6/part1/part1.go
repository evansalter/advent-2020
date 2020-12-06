package part1

import (
	"fmt"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

func Run() {
	lines := helpers.ReadInputFile(6)
	groups := make([]map[string]interface{}, 0)
	groupAnswers := make(map[string]interface{}, 0)
	for _, l := range lines {
		if l == "" {
			groups = append(groups, groupAnswers)
			groupAnswers = make(map[string]interface{}, 0)
		}
		questions := strings.Split(l, "")
		for _, q := range questions {
			groupAnswers[q] = struct{}{}
		}
	}

	sum := 0
	for _, g := range groups {
		sum += len(g)
	}

	fmt.Println(sum)
}
