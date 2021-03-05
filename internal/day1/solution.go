package day1

import (
	"fmt"
	"strconv"
	"strings"

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
func (p *Puzzle) Load() error {
	rawData, err := common.LoadPuzzleInput(dayNumber)
	if err != nil {
		return fmt.Errorf("failed to load input: %s", err)
	}

	for _, strVal := range strings.Split(string(rawData), "\n") {
		if strVal == "" {
			// Blank line at the end of file
			break
		}
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			return fmt.Errorf("error while reading input: %s", err)
		}
		p.expenses = append(p.expenses, intVal)
	}
	return nil
}

// Reload does nothing since it is not necessary for day 1.
func (p *Puzzle) Reload() error {
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
	return "", fmt.Errorf("failed to solve part 1 of day %d's puzzle", dayNumber)
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
	return "", fmt.Errorf("failed to solve part 1 of day %d's puzzle", dayNumber)
}
