package day9

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 9

	defaultBlockSize = 25
)

// Puzzle contains the puzzle for day 9.
type Puzzle struct {
	blockSize     int
	data          []int
	preamable     map[int]struct{}
	invalidNumber int
}

// Day returns 9.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 9.
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

	return nil
}

// SolvePart1 returns the answer to day 9, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	if p.blockSize == 0 {
		p.blockSize = defaultBlockSize
	}

	p.buildPreamable(p.data[:p.blockSize]...)
forLoop:
	for i, iEnd := p.blockSize, len(p.data); i < iEnd; i++ {
		sum := p.data[i]
		for v1 := range p.preamable {
			v2 := sum - v1
			if v2 != v1 && p.lookupPreamable(v2) {
				p.updateNewPreamable(i)
				continue forLoop
			}
		}

		// Could not find the subcomponents
		p.invalidNumber = sum
		return common.ToAnswer(p.invalidNumber), nil
	}

	return "", errors.New("could not find the answer")
}

func (p *Puzzle) buildPreamable(values ...int) {
	p.preamable = make(map[int]struct{})
	for _, v := range values {
		p.preamable[v] = struct{}{}
	}
}

func (p *Puzzle) updateNewPreamable(idx int) {
	delete(p.preamable, p.data[idx-p.blockSize])
	p.preamable[p.data[idx]] = struct{}{}
}

func (p *Puzzle) lookupPreamable(value int) bool {
	_, ok := p.preamable[value]
	return ok
}

// SolvePart2 returns the answer to day 9, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	if p.blockSize == 0 {
		p.blockSize = defaultBlockSize
	}

	l, r := 0, 0
	s := 0 // s is the sum of p.data[l:r]
	for r < len(p.data) {
		if s == p.invalidNumber {
			min, max := p.findMinMax(p.data[l:r])
			return common.ToAnswer(min + max), nil
		} else if s < p.invalidNumber {
			s += p.data[r]
			r++
		} else {
			s -= p.data[l]
			l++
		}
	}
	return "", errors.New("could not find the answer")
}

func (p *Puzzle) findMinMax(values []int) (int, int) {
	min, max := 1<<31, 0
	for _, v := range values {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}
