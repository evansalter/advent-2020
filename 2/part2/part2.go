package part2

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

		idxA, idxB, lttr, psswd := parseLine(l)
		if ok := checkPassword(idxA, idxB, lttr, psswd); ok {
			numCorrect++
		}
	}

	fmt.Println(numCorrect)
}

func parseLine(l string) (idxA, idxB int, lttr, psswd string) {
	s := strings.Split(l, " ")
	lttr, psswd = strings.Trim(s[1], ":"), s[2]
	s = strings.Split(s[0], "-")
	minStr, maxStr := s[0], s[1]
	var err error
	idxA, err = strconv.Atoi(minStr)
	if err != nil {
		panic(err)
	}
	idxB, err = strconv.Atoi(maxStr)
	if err != nil {
		panic(err)
	}
	return
}

func checkPassword(idxA, idxB int, lttr, psswd string) bool {
	numTimesLetterFound := 0

	letters := strings.Split(psswd, "")
	charsToCheck := []string{letters[idxA-1], letters[idxB-1]}
	for _, char := range charsToCheck {
		if char == lttr {
			numTimesLetterFound++
		}
	}

	return numTimesLetterFound == 1
}
