package day1

import (
	"context"
	"fmt"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 1

	requiredSum = 2020
)

// Puzzle contains the puzzle for day 1.
type Puzzle struct {
	expenses []int
}

// Day returns 1.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 1.
func (p *Puzzle) Load(ctx context.Context) error {
	datastream := common.LoadInputAsync(ctx, p.Day(), common.ChannelSizeDefault)
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		intVal, err := strconv.Atoi(string(dataline.Content))
		if err != nil {
			return fmt.Errorf("invalid line input: %s", err)
		}
		p.expenses = append(p.expenses, intVal)
	}

	return nil
}

// Reload is not necessary for day 1.
func (p *Puzzle) Reload(context.Context) error {
	return nil
}

// SolvePart1 returns the answer to day 1, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	lookup := make(map[int]struct{})
	for _, v1 := range p.expenses {
		v2 := requiredSum - v1
		if _, ok := lookup[v2]; ok {
			return common.ToAnswer(v1 * v2), nil
		}

		lookup[v1] = struct{}{}
	}
	return "", common.ErrCannotComputeAnswer
}

// SolvePart2 returns the answer to day 1, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	lookup := make(map[int]int)
	n := len(p.expenses)
	for i, iEnd := 0, n-1; i < iEnd; i++ {
		for j := i + 1; j < n; j++ {
			v1, v2 := p.expenses[i], p.expenses[j]
			lookup[v1+v2] = v1 * v2
		}
	}

	for _, v3 := range p.expenses {
		if v12, ok := lookup[requiredSum-v3]; ok {
			return common.ToAnswer(v12 * v3), nil
		}
	}
	return "", common.ErrCannotComputeAnswer
}
