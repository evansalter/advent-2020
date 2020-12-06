package part2

import (
	"fmt"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

type GroupAnswers struct {
	NumPeople int
	Answers   [][]string
}

func (g *GroupAnswers) GetNumAnswered() map[string]int {
	m := make(map[string]int)
	for _, person := range g.Answers {
		for _, q := range person {
			m[q]++
		}
	}
	return m
}

func (g *GroupAnswers) GetNumAllAnswered() (num int) {
	for _, count := range g.GetNumAnswered() {
		if count == g.NumPeople {
			num++
		}
	}
	return
}

func Run() {
	lines := helpers.ReadInputFile(6)
	var groupAnswers []*GroupAnswers
	g := &GroupAnswers{}
	for _, l := range lines {
		if l == "" {
			groupAnswers = append(groupAnswers, g)
			g = &GroupAnswers{}
			continue
		}
		g.NumPeople++
		g.Answers = append(g.Answers, strings.Split(l, ""))
	}

	sum := 0
	for _, g := range groupAnswers {
		sum += g.GetNumAllAnswered()
	}

	fmt.Println(sum)
}
