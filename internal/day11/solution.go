package day11

import (
	"context"
	"errors"
	"fmt"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 11

	statusEmpty    = 'L'
	statusOccupied = '#'
	statusFloor    = '.'
)

// Puzzle contains the puzzle for day 11.
type Puzzle struct {
	original [][]rune
	data     [][]rune
	nrow     int
	ncol     int
}

// Day returns 11.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 11.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		p.original = append(p.original, []rune(string(dataline.Content)))
	}

	if len(p.original) == 0 {
		return errors.New("empty input data")
	}
	p.nrow, p.ncol = len(p.original), len(p.original[0])
	for _, row := range p.original {
		if len(row) != p.ncol {
			return fmt.Errorf("column length mismatched for row \"%s\": expected %d, got %d", string(row), p.ncol, len(row))
		}
	}
	return nil
}

// SolvePart1 returns the answer to day 11, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	p.data = p.original
	for {
		hasChanges := p.updateOnce(p.getNewStatePart1)
		if !hasChanges {
			return common.ToAnswer(p.countOccupiedSeats()), nil
		}
	}
}

// SolvePart2 returns the answer to day 11, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	p.data = p.original
	for {
		hasChanges := p.updateOnce(p.getNewStatePart2)
		if !hasChanges {
			return common.ToAnswer(p.countOccupiedSeats()), nil
		}
	}
}

// updateOnce update the entire seat layout once. It returns true if at least
// a seat's status is changed.
func (p *Puzzle) updateOnce(updateRule func(int, int) rune) bool {
	hasChanges := false
	nextdata := make([][]rune, p.nrow)
	for r := 0; r < p.nrow; r++ {
		nextdata[r] = make([]rune, p.ncol)
		for c := 0; c < p.ncol; c++ {
			nextdata[r][c] = updateRule(r, c)
			if nextdata[r][c] != p.data[r][c] {
				hasChanges = true
			}
		}
	}
	p.data = nextdata
	return hasChanges
}

func (p *Puzzle) getNewStatePart1(r, c int) rune {
	switch p.data[r][c] {
	case statusEmpty:
		// return occupied if all adjacent seats are not
		if p.countAdjacentOccupiedSeats(r, c) == 0 {
			return statusOccupied
		}
	case statusOccupied:
		// return empty if 4 or more adjacent seats are occupied
		if p.countAdjacentOccupiedSeats(r, c) >= 4 {
			return statusEmpty
		}
	}
	return p.data[r][c] // unchanged
}

func (p *Puzzle) countAdjacentOccupiedSeats(r, c int) int {
	count := 0
	for _, v := range [][]int{
		{r - 1, c - 1},
		{r - 1, c},
		{r - 1, c + 1},
		{r, c - 1},
		{r, c + 1},
		{r + 1, c - 1},
		{r + 1, c},
		{r + 1, c + 1},
	} {
		if p.isValid(v[0], v[1]) && p.data[v[0]][v[1]] == statusOccupied {
			count++
		}
	}
	return count
}

func (p *Puzzle) getNewStatePart2(r, c int) rune {
	switch p.data[r][c] {
	case statusEmpty:
		// return occupied if all directional seats are not
		if p.countDirectionalOccupiedSeats(r, c) == 0 {
			return statusOccupied
		}
	case statusOccupied:
		// return empty if 5 or more directional seats are occupied
		if p.countDirectionalOccupiedSeats(r, c) >= 5 {
			return statusEmpty
		}
	}
	return p.data[r][c] // unchanged
}

func (p *Puzzle) countDirectionalOccupiedSeats(r, c int) int {
	count := 0
	if p.hasOccupiedSeatsInDirection(r, c, -1, -1) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, -1, 0) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, -1, 1) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, 0, -1) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, 0, 1) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, 1, -1) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, 1, 0) {
		count++
	}
	if p.hasOccupiedSeatsInDirection(r, c, 1, 1) {
		count++
	}
	return count
}

func (p *Puzzle) hasOccupiedSeatsInDirection(r, c, dr, dc int) bool {
	for i := 1; ; i++ {
		rx, cx := r+dr*i, c+dc*i
		if !p.isValid(rx, cx) {
			return false
		}
		if p.data[rx][cx] != statusFloor {
			return p.data[rx][cx] == statusOccupied
		}
	}
}

func (p *Puzzle) isValid(r, c int) bool {
	return r >= 0 && c >= 0 && r < p.nrow && c < p.ncol
}

func (p *Puzzle) countOccupiedSeats() int {
	count := 0
	for _, row := range p.data {
		for _, v := range row {
			if v == statusOccupied {
				count++
			}
		}
	}
	return count
}
