package common

import (
	"context"
	"fmt"
	"time"
)

// Puzzle represents a specific day's puzzle.
type Puzzle interface {
	// Day returns day number of the puzzle
	Day() int

	// Load loads the puzzle input.
	Load(context.Context, <-chan LineContent) error

	// SolvePart1 solves and returns part 1's answer to the puzzle.
	SolvePart1() (Answer, error)

	// SolvePart2 solves and returns part 2's answer to the puzzle.
	SolvePart2() (Answer, error)
}

const (
	// ChannelSizeDefault is the default size for the channels created by functions
	// in this package.
	ChannelSizeDefault = 5

	// AlgorithmTimeout is the time out duration to solve a day's puzzle.
	AlgorithmTimeout = 30 * time.Second
)

// Solve returns the answer to the puzzle.
func Solve(p Puzzle) (*AnswerOfDay, error) {
	ctx, cancel := context.WithTimeout(context.Background(), AlgorithmTimeout)
	defer cancel()

	datastream := loadInputAsync(ctx, p.Day(), ChannelSizeDefault)
	err := p.Load(ctx, datastream)
	if err == nil {
		err = ctx.Err()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %s", err)
	}

	return computeAnswer(p)
}

// SolveExample returns the answer to the example given in the puzzle.
func SolveExample(p Puzzle) (*AnswerOfDay, error) {
	ctx, cancel := context.WithTimeout(context.Background(), AlgorithmTimeout)
	defer cancel()

	datastream := loadExampleInputAsync(ctx, p.Day(), ChannelSizeDefault)
	err := p.Load(ctx, datastream)
	if err == nil {
		err = ctx.Err()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %s", err)
	}

	return computeAnswer(p)
}

func computeAnswer(p Puzzle) (*AnswerOfDay, error) {
	ans1, err := p.SolvePart1()
	if err != nil {
		return nil, fmt.Errorf("failed to solve part 1: %s", err)
	}
	ans2, err := p.SolvePart2()
	if err != nil {
		return nil, fmt.Errorf("failed to solve part 2: %s", err)
	}
	return &AnswerOfDay{
		Day:    p.Day(),
		First:  ans1,
		Second: ans2,
	}, nil
}
