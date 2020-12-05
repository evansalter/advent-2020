package part1

import (
	"fmt"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p *Passport) Set(key, value string) {
	switch key {
	case "byr":
		p.BirthYear = value
	case "iyr":
		p.IssueYear = value
	case "eyr":
		p.ExpirationYear = value
	case "hgt":
		p.Height = value
	case "hcl":
		p.HairColor = value
	case "ecl":
		p.EyeColor = value
	case "pid":
		p.PassportID = value
	case "cid":
		p.CountryID = value
	}
}

func (p *Passport) IsValid() bool {
	if p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != "" {
		return true
	}
	return false
}

func Run() {
	lines := helpers.ReadInputFile(4)
	allPassports := make([]*Passport, 0)
	passportLines := make([]string, 0)
	for _, l := range lines {
		if l == "" {
			p := passportFromLines(passportLines)
			allPassports = append(allPassports, p)
			passportLines = make([]string, 0)
		}
		passportLines = append(passportLines, l)
	}
	numValid := 0
	for _, p := range allPassports {
		if p.IsValid() {
			numValid++
		}
	}
	fmt.Println(numValid)
}

func passportFromLines(lines []string) *Passport {
	p := &Passport{}
	for _, l := range lines {
		parts := strings.Split(l, " ")
		for _, part := range parts {
			if part == "" {
				continue
			}
			keyValue := strings.Split(part, ":")
			key, value := keyValue[0], keyValue[1]
			p.Set(key, value)
		}
	}
	return p
}
