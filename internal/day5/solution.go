package day5

import (
	"context"
	"fmt"
	"strings"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 5

	codeFront = 'F'
	codeBack  = 'B'
	codeLeft  = 'L'
	codeRight = 'R'
)

type seatCode []rune

func (sc *seatCode) decode() (*seatPosition, error) {
	scdata := []rune(*sc)
	if len(scdata) != 10 {
		return nil, fmt.Errorf("seat code \"%s\": invalid code length: %d", string(scdata), len(scdata))
	}
	sp := &seatPosition{0, 0}
	for i := 0; i < 7; i++ {
		if scdata[i] == codeBack {
			sp.row += 1 << (6 - i)
		} else if scdata[i] != codeFront {
			return nil, fmt.Errorf("invalid character '%c' in seat code at position %d", scdata[i], i)
		}
	}

	for i := 7; i < 10; i++ {
		if scdata[i] == codeRight {
			sp.col += 1 << (9 - i)
		} else if scdata[i] != codeLeft {
			return nil, fmt.Errorf("invalid character '%c' in seat code at position %d", scdata[i], i)
		}
	}
	return sp, nil
}

type seatPosition struct {
	row int
	col int
}

func (sp *seatPosition) ID() int {
	return sp.row*8 + sp.col
}

// Puzzle contains the puzzle for day 5.
type Puzzle struct {
	spList []seatPosition
}

// Day returns 5.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 5.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		sc := seatCode(strings.TrimSpace(string(dataline.Content)))
		sp, err := sc.decode()
		if err != nil {
			return err
		}
		p.spList = append(p.spList, *sp)
	}

	return nil
}

// SolvePart1 returns the answer to day 5, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	highestID := 0
	for _, sp := range p.spList {
		id := sp.ID()
		if id > highestID {
			highestID = id
		}
	}

	return common.ToAnswer(highestID), nil
}

// SolvePart2 returns the answer to day 5, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	const ncols = 8
	sortSeatPosition(p.spList)
	myseat := &seatPosition{p.spList[0].row, p.spList[0].col}
	for _, sp := range p.spList {
		if myseat.row != sp.row || myseat.col != sp.col {
			fmt.Println(myseat)
			return common.ToAnswer(myseat.ID()), nil
		}

		myseat.col++
		if myseat.col == ncols {
			myseat.row++
			myseat.col = 0
		}
	}
	return common.ToAnswer(myseat.ID()), nil
}
