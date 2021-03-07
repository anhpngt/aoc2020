package day3

import (
	"context"
	"errors"
	"fmt"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 3

	squareOpen = '.'
	squareTree = '#'
)

type grid struct {
	nrows int
	ncols int
	data  [][]rune
}

// moveDelta computes the next (r, c) position after moving from (r, c) for (dr, dc).
// It wraps at the border of the map.
func (g *grid) moveDelta(r, c, dr, dc int) (int, int) {
	r += dr
	if r >= g.nrows {
		r -= g.nrows
	}
	c += dc
	if c >= g.ncols {
		c -= g.ncols
	}
	return r, c
}

// at returns the value at (r, c) on the map.
func (g *grid) at(r, c int) rune {
	return g.data[r][c]
}

// newGrid verifies data and returns a new grid.
func newGrid(nrows, ncols int, data [][]rune) (*grid, error) {
	if nrows != len(data) {
		return nil, fmt.Errorf("grid data size mismatched: expecting %d rows, got %d rows", nrows, len(data))
	}
	for r := 0; r < nrows; r++ {
		if ncols != len(data[r]) {
			return nil, fmt.Errorf(
				"grid data size mismatched: expecting %d columns, got %d columns for \"%s\"",
				ncols, len(data[r]), string(data[r]),
			)
		}
	}
	return &grid{nrows, ncols, data}, nil
}

// Puzzle contains the puzzle for day 3.
type Puzzle struct {
	travelMap *grid
}

// Day returns 3.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 3.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan common.LineContent) error {
	gridData := make([][]rune, 0)
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		gridData = append(gridData, []rune(string(dataline.Content)))
	}

	if len(gridData) == 0 {
		return errors.New("the input data is empty")
	}
	var err error
	p.travelMap, err = newGrid(len(gridData), len(gridData[0]), gridData)
	return err
}

// countTreesEncountered counts the number of tree encountered if slope (dr, dc) is chosen.
func (p *Puzzle) countTreesEncountered(dr, dc int) int {
	r, c := 0, 0
	count := 0
	for r < p.travelMap.nrows-1 {
		r, c = p.travelMap.moveDelta(r, c, dr, dc)
		if p.travelMap.at(r, c) == squareTree {
			count++
		}
	}
	return count
}

// SolvePart1 returns the answer to day 3, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	count := p.countTreesEncountered(1, 3)
	return common.ToAnswer(count), nil
}

// SolvePart2 returns the answer to day 3, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	drList := []int{1, 1, 1, 1, 2}
	dcList := []int{1, 3, 5, 7, 1}
	countProd := 1
	for i := 0; i < len(drList); i++ {
		count := p.countTreesEncountered(drList[i], dcList[i])
		countProd *= count
	}
	return common.ToAnswer(countProd), nil
}
