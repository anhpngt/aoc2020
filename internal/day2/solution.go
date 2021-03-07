package day2

import (
	"context"
	"fmt"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 2
)

type password []rune

type passwordPolicy struct {
	LowerLimit int
	UpperLimit int
	Letter     rune
}

func (pwp *passwordPolicy) validateCount(pw password) bool {
	count := 0
	for _, c := range pw {
		if c == pwp.Letter {
			count++
		}
	}
	return pwp.LowerLimit <= count && count <= pwp.UpperLimit
}

func (pwp *passwordPolicy) validatePosition(pw password) bool {
	count := 0
	if pwp.LowerLimit-1 < len(pw) && pw[pwp.LowerLimit-1] == pwp.Letter {
		count++
	}
	if pwp.UpperLimit-1 < len(pw) && pw[pwp.UpperLimit-1] == pwp.Letter {
		count++
	}
	return count == 1
}

// Puzzle contains the puzzle for day 2.
type Puzzle struct {
	passwordList       []password
	passwordPolicyList []passwordPolicy
}

// Day returns 2.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 2.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan common.LineContent) error {
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		var pwstr string
		pwp := passwordPolicy{}
		if _, err := fmt.Sscanf(
			string(dataline.Content),
			"%d-%d %c: %s",
			&pwp.LowerLimit, &pwp.UpperLimit, &pwp.Letter, &pwstr,
		); err != nil {
			return fmt.Errorf("invalid line input: \"%s\", error: %s", string(dataline.Content), err)
		}

		p.passwordList = append(p.passwordList, password(pwstr))
		p.passwordPolicyList = append(p.passwordPolicyList, pwp)
	}
	return nil
}

// SolvePart1 returns the answer to day 2, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	count := 0
	for i := range p.passwordList {
		if p.passwordPolicyList[i].validateCount(p.passwordList[i]) {
			count++
		}
	}
	return common.ToAnswer(count), nil
}

// SolvePart2 returns the answer to day 2, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	count := 0
	for i := range p.passwordList {
		if p.passwordPolicyList[i].validatePosition(p.passwordList[i]) {
			count++
		}
	}
	return common.ToAnswer(count), nil
}
