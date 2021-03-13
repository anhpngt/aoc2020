package day10

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 10
)

// Puzzle contains the puzzle for day 10.
type Puzzle struct {
	data []int
}

// Day returns 10.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 10.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}
		intVal, err := strconv.Atoi(string(dataline.Content))
		if err != nil {
			return fmt.Errorf("invalid data line: %s", err)
		}
		p.data = append(p.data, intVal)
	}

	p.data = append(p.data, 0) // add value for the charging outlet
	sort.Ints(p.data)
	return nil
}

// SolvePart1 returns the answer to day 10, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	diffone, diffthree := 0, 1 // our device is always +3
	for i, iEnd := 1, len(p.data); i < iEnd; i++ {
		diff := p.data[i] - p.data[i-1]
		if diff == 1 {
			diffone++
		} else if diff == 3 {
			diffthree++
		}
	}
	return common.ToAnswer(diffone * diffthree), nil
}

// SolvePart2 returns the answer to day 10, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	counts := make([]int64, len(p.data))
	counts[0] = 1
	for i, iEnd := 1, len(counts); i < iEnd; i++ {
		for j := i - 1; j >= 0 && p.data[i]-p.data[j] <= 3; j-- {
			counts[i] += counts[j]
		}
	}

	return common.ToAnswerInt64(counts[len(counts)-1]), nil
}
