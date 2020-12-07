package part2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/evansalter/advent-2020/helpers"
)

const (
	TargetBagColor = "shiny gold"
)

type Bag struct {
	BagColor string
	Num      int
}

type BagRule struct {
	BagColor string
	Contains []*Bag
}

func Run() {
	lines := helpers.ReadInputFile(7)
	allRules := make([]*BagRule, len(lines)-1)
	for i, l := range lines {
		if l == "" {
			continue
		}
		r := regexp.MustCompile("^([a-z ]+) bags contain")
		br := &BagRule{
			BagColor: r.FindStringSubmatch(l)[1],
		}

		r = regexp.MustCompile("(\\d+) ([a-z ]+) bag")
		contains := make([]*Bag, 0)
		for _, m := range r.FindAllStringSubmatch(l, -1) {
			count, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}
			contains = append(contains, &Bag{
				BagColor: m[2],
				Num:      count,
			})
		}
		br.Contains = contains
		allRules[i] = br
	}

	count := CountNumBags(allRules, TargetBagColor)
	fmt.Println(count)
}

func CountNumBags(allRules []*BagRule, target string) int {
	count := 0
	for _, r := range allRules {
		if r.BagColor == target {
			for _, c := range r.Contains {
				count += c.Num
				count += c.Num * CountNumBags(allRules, c.BagColor)
			}
			break
		}
	}
	return count
}

//func FindMatchingRules(allRules []*BagRule, target string) []*BagRule {
//	matchingRules := make([]*BagRule, 0)
//	for _, r := range allRules {
//		for _, c := range r.Contains {
//			if c.BagColor == target {
//				matchingRules = append(matchingRules, r)
//				matchingRules = append(matchingRules, FindMatchingRules(allRules, r.BagColor)...)
//			}
//		}
//	}
//	return matchingRules
//}
