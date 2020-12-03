package part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

func Run() {
	lines := helpers.ReadInputFile(2)
	numCorrect := 0
	for _, l := range lines {
		if l == "" {
			continue
		}

		min, max, lttr, psswd := parseLine(l)
		if ok := checkPassword(min, max, lttr, psswd); ok {
			numCorrect++
		}
	}

	fmt.Println(numCorrect)
}

func parseLine(l string) (min, max int, lttr, psswd string) {
	s := strings.Split(l, " ")
	lttr, psswd = strings.Trim(s[1], ":"), s[2]
	s = strings.Split(s[0], "-")
	minStr, maxStr := s[0], s[1]
	var err error
	min, err = strconv.Atoi(minStr)
	if err != nil {
		panic(err)
	}
	max, err = strconv.Atoi(maxStr)
	if err != nil {
		panic(err)
	}
	return
}

func checkPassword(min, max int, lttr, psswd string) bool {
	numTimesLetterFound := 0

	for _, l := range strings.Split(psswd, "") {
		if l == lttr {
			numTimesLetterFound++
		}
	}

	return numTimesLetterFound >= min && numTimesLetterFound <= max
}
