package part2

import (
	"fmt"
	"regexp"
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

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//hgt (Height) - a number followed by either cm or in:
//	If cm, the number must be at least 150 and at most 193.
//	If in, the number must be at least 59 and at most 76.
//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//pid (Passport ID) - a nine-digit number, including leading zeroes.
//cid (Country ID) - ignored, missing or not.

func (p *Passport) IsValid() bool {
	if p.BirthYear == "" ||
		p.IssueYear == "" ||
		p.ExpirationYear == "" ||
		p.Height == "" ||
		p.HairColor == "" ||
		p.EyeColor == "" ||
		p.PassportID == "" {
		return false
	}
	if p.BirthYear < "1920" || p.BirthYear > "2002" {
		fmt.Printf("BirthYear not valid: %s\n", p.BirthYear)
		return false
	}
	if p.IssueYear < "2010" || p.IssueYear > "2020" {
		fmt.Printf("IssueYear not valid: %s\n", p.IssueYear)
		return false
	}
	if p.ExpirationYear < "2020" || p.ExpirationYear > "2030" {
		fmt.Printf("ExpirationYear not valid: %s\n", p.ExpirationYear)
		return false
	}
	if strings.HasSuffix(p.Height, "cm") {
		h := strings.Trim(p.Height, "cm")
		if h < "150" || h > "193" {
			fmt.Printf("Height not valid: %s\n", p.Height)
			return false
		}
	} else if strings.HasSuffix(p.Height, "in") {
		h := strings.Trim(p.Height, "in")
		if h < "59" || h > "76" {
			fmt.Printf("Height not valid: %s\n", p.Height)
			return false
		}
	} else {
		fmt.Printf("Height not valid: %s\n", p.Height)
		return false
	}
	m, _ := regexp.MatchString("^#[0-9a-f]{6}$", p.HairColor)
	if !m {
		fmt.Printf("HairColor not valid: %s\n", p.HairColor)
		return false
	}
	if !stringOneOf(p.EyeColor, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) {
		fmt.Printf("EyeColor not valid: %s\n", p.EyeColor)
		return false
	}
	m, _ = regexp.MatchString("^[0-9]{9}$", p.PassportID)
	if !m {
		fmt.Printf("PassportID not valid: %s\n", p.PassportID)
		return false
	}
	return true
}

func stringOneOf(s string, options []string) bool {
	for _, o := range options {
		if s == o {
			return true
		}
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
