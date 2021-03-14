package day12

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 12
)

var insRe = regexp.MustCompile(`^(?:(N|S|E|W|F)(\d+))|(?:(L|R)(90|180|270))$`)

type instruction struct {
	action rune
	value  int
}

func parseInstruction(s []byte) (*instruction, error) {
	matches := insRe.FindSubmatch(s)
	if matches == nil {
		return nil, errors.New("invalid format")
	}

	actionMatch := matches[1]
	valueMatch := matches[2]
	if len(actionMatch) == 0 {
		actionMatch = matches[3]
		valueMatch = matches[4]
	}

	var action rune
	for _, c := range string(actionMatch) {
		action = c
		break
	}
	value, _ := strconv.Atoi(string(valueMatch))
	return &instruction{action, value}, nil
}

// Puzzle contains the puzzle for day 12.
type Puzzle struct {
	insList []instruction
}

// Day returns 12.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 12.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		ins, err := parseInstruction(dataline.Content)
		if err != nil {
			return fmt.Errorf("cannot parse instruction \"%s\": %s", string(dataline.Content), err)
		}
		p.insList = append(p.insList, *ins)
	}

	return nil
}

// SolvePart1 returns the answer to day 12, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	pos := initialPosition()
	for _, ins := range p.insList {
		pos.move(ins.action, ins.value)
	}
	return common.ToAnswer(pos.mhtDist()), nil
}

// SolvePart2 returns the answer to day 12, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	pos := initialPositionWithWaypoint()
	for _, ins := range p.insList {
		pos.move(ins.action, ins.value)
	}
	return common.ToAnswer(pos.pos.mhtDist()), nil
}

func initialPosition() position {
	return position{0, 0, heading('E')}
}

func initialPositionWithWaypoint() positionWithWaypoint {
	return positionWithWaypoint{
		pos: position{0, 0, heading('E')},
		wp:  position{10, 1, heading('E')},
	}
}
