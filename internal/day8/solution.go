package day8

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 8
)

type opType string

const (
	acc opType = "acc"
	jmp opType = "jmp"
	nop opType = "nop"
)

var opRe = regexp.MustCompile(`^(acc|jmp|nop) ((?:\+|-)\d+)$`)

type op struct {
	optype opType
	opval  int
}

func newOp(b []byte) (*op, error) {
	matches := opRe.FindSubmatch(b)
	if matches == nil {
		return nil, errors.New("invalid format")
	}

	o := &op{}
	switch string(matches[1]) {
	case "acc":
		o.optype = acc
	case "jmp":
		o.optype = jmp
	default: // "nop"
		o.optype = nop
	}

	o.opval, _ = strconv.Atoi(string(matches[2]))
	return o, nil
}

// Puzzle contains the puzzle for day 8.
type Puzzle struct {
	instructions []*op
	pc           int
	accVal       int
}

// Day returns 8.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 8.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		o, err := newOp(dataline.Content)
		if err != nil {
			return fmt.Errorf(`failed to parse instruction "%s": %s`, string(dataline.Content), err)
		}

		p.instructions = append(p.instructions, o)
	}

	return nil
}

// run executes the instruction set and returns true if there is no infinite loop.
func (p *Puzzle) run() bool {
	p.pc = 0
	p.accVal = 0
	visited := make(map[int]struct{})
	for {
		o := p.instructions[p.pc]
		visited[p.pc] = struct{}{}
		switch o.optype {
		case acc:
			p.accVal += o.opval
			p.pc++
		case jmp:
			p.pc += o.opval
		default: // nop
			p.pc++
		}

		if p.pc >= len(p.instructions) {
			return true
		}
		if _, yes := visited[p.pc]; yes {
			return false
		}
	}
}

// SolvePart1 returns the answer to day 8, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	_ = p.run()
	return common.ToAnswer(p.accVal), nil
}

// SolvePart2 returns the answer to day 8, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
loop:
	for _, o := range p.instructions {
		switch o.optype {
		case jmp:
			o.optype = nop
			if p.run() {
				break loop
			}
			o.optype = jmp
		case nop:
			o.optype = jmp
			if p.run() {
				break loop
			}
			o.optype = nop
		}
	}
	return common.ToAnswer(p.accVal), nil
}
