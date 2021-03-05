package common

import "fmt"

// Puzzle represents a specific day's puzzle.
type Puzzle interface {
	// Day returns day number of the puzzle
	Day() int

	// Load loads the puzzle input.
	Load() error

	// Reload reloads the puzzle input, if needed, after solving part 1.
	Reload() error

	// SolvePart1 solves and returns part 1's answer to the puzzle.
	SolvePart1() (Answer, error)

	// SolvePart2 solves and returns part 2's answer to the puzzle.
	SolvePart2() (Answer, error)
}

// Solve returns the answer to the puzzle.
func Solve(p Puzzle) (*AnswerOfDay, error) {
	if err := p.Load(); err != nil {
		return nil, fmt.Errorf("cannot load puzzle for day %d: %s", p.Day(), err)
	}
	ans1, err := p.SolvePart1()
	if err != nil {
		return nil, fmt.Errorf("cannot solve part 1 of day %d: %s", p.Day(), err)
	}

	if err = p.Reload(); err != nil {
		return nil, fmt.Errorf("cannot reload puzzle for day %d: %s", p.Day(), err)
	}
	ans2, err := p.SolvePart2()
	if err != nil {
		return nil, fmt.Errorf("cannot solve part 2 of day %d: %s", p.Day(), err)
	}
	return &AnswerOfDay{
		Day:    p.Day(),
		First:  ans1,
		Second: ans2,
	}, nil
}
