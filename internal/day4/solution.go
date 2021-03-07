package day4

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/anhpngt/aoc2020/internal/common"
)

const dayNumber = 4

var (
	byrRe = regexp.MustCompile(`^((19[2-9][0-9])|(200[0-2]))$`)
	iyrRe = regexp.MustCompile(`^20(1[0-9]|20)$`)
	eyrRe = regexp.MustCompile(`^20(2[0-9]|30)$`)
	hgtRe = regexp.MustCompile(`^(((1[5-8][0-9]|19[0-3])cm)|((59|6\d|7[0-6])in))$`)
	hclRe = regexp.MustCompile(`^#[a-f0-9]{6}$`)
	eclRe = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRe = regexp.MustCompile(`^\d{9}$`)
)

// passport represents the passport in the puzzle.
type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) mapKeyValue(s string) error {
	splits := strings.Split(s, ":")
	if len(splits) != 2 {
		return fmt.Errorf("invalid format for key:value pair: %s", s)
	}
	*p.lookup(splits[0]) = splits[1]
	return nil
}

func (p *passport) lookup(key string) *string {
	switch key {
	case "byr":
		return &p.byr
	case "iyr":
		return &p.iyr
	case "eyr":
		return &p.eyr
	case "hgt":
		return &p.hgt
	case "hcl":
		return &p.hcl
	case "ecl":
		return &p.ecl
	case "pid":
		return &p.pid
	case "cid":
		return &p.cid
	default:
		panic(fmt.Errorf("invalid key in key:value pair: %s", key))
	}
}

// isRequiredFieldsPresent returns true if the passport has all fields possibly except cid.
func (p *passport) isRequiredFieldsPresent() bool {
	return len(p.byr) > 0 &&
		len(p.iyr) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0
}

func (p *passport) isRequiredFieldsValid() bool {
	return p.isBYRValid() &&
		p.isIYRValid() &&
		p.isEYRValid() &&
		p.isHGTValid() &&
		p.isHCLValid() &&
		p.isECLValid() &&
		p.isPIDValid()
}

func (p *passport) isBYRValid() bool {
	return byrRe.FindStringIndex(p.byr) != nil
}

func (p *passport) isIYRValid() bool {
	return iyrRe.FindStringIndex(p.iyr) != nil
}

func (p *passport) isEYRValid() bool {
	return eyrRe.FindStringIndex(p.eyr) != nil
}

func (p *passport) isHGTValid() bool {
	return hgtRe.FindStringIndex(p.hgt) != nil
}

func (p *passport) isHCLValid() bool {
	return hclRe.FindStringIndex(p.hcl) != nil
}

func (p *passport) isECLValid() bool {
	return eclRe.FindStringIndex(p.ecl) != nil
}

func (p *passport) isPIDValid() bool {
	return pidRe.FindStringIndex(p.pid) != nil
}

// Puzzle contains the puzzle for day 4.
type Puzzle struct {
	passportList []*passport
}

// Day returns 4.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 4.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	var pprt *passport
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		// Initialization
		if pprt == nil {
			pprt = &passport{}
		}

		// Blank line signals a new passport
		linestr := strings.TrimSpace(string(dataline.Content))
		if len(linestr) == 0 {
			p.passportList = append(p.passportList, pprt)
			pprt = &passport{}
			continue
		}

		kvList := strings.Split(linestr, " ")
		for _, kv := range kvList {
			if err := pprt.mapKeyValue(kv); err != nil {
				return err
			}
		}
	}

	if pprt != nil {
		p.passportList = append(p.passportList, pprt)
	}
	return nil
}

// SolvePart1 returns the answer to day 4, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	count := 0
	for _, pprt := range p.passportList {
		if pprt.isRequiredFieldsPresent() {
			count++
		}
	}
	return common.ToAnswer(count), nil
}

// SolvePart2 returns the answer to day 4, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	count := 0
	for _, pprt := range p.passportList {
		if pprt.isRequiredFieldsPresent() && pprt.isRequiredFieldsValid() {
			count++
		}
	}
	return common.ToAnswer(count), nil
}
