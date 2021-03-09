package day6

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 6
)

type mcq struct {
	answer    [26]int
	groupsize int
}

// addFrom adds answers from s to mcq.
func (m *mcq) addFrom(s string) error {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return fmt.Errorf("invalid character \"%c\" in yes-question string", c)
		}
		m.answer[c-'a']++
	}
	m.groupsize++
	return nil
}

// countOR returns the number of questions to which at least one person from the group
// answers 'yes'.
func (m *mcq) countOR() int {
	count := 0
	for _, v := range m.answer {
		if v > 0 {
			count++
		}
	}
	return count
}

// countAND returns the number of questions to which everyone from the group answers 'yes'.
func (m *mcq) countAND() int {
	count := 0
	for _, v := range m.answer {
		if v == m.groupsize {
			count++
		}
	}
	return count
}

// Puzzle contains the puzzle for day 6.
type Puzzle struct {
	answerList []mcq
}

// Day returns 6.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 6.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	var m *mcq
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		s := strings.TrimSpace(string(dataline.Content))
		if len(s) == 0 {
			if m == nil {
				return errors.New("unexpected blank line in input file")
			}
			p.answerList = append(p.answerList, *m)
			m = nil
			continue
		} else {
			if m == nil {
				m = &mcq{}
			}
			m.addFrom(s)
		}
	}

	if m != nil {
		p.answerList = append(p.answerList, *m)
	}

	return nil
}

// SolvePart1 returns the answer to day 6, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	count := 0
	for _, m := range p.answerList {
		count += m.countOR()
	}
	return common.ToAnswer(count), nil
}

// SolvePart2 returns the answer to day 6, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	count := 0
	for _, m := range p.answerList {
		count += m.countAND()
	}
	return common.ToAnswer(count), nil
}
