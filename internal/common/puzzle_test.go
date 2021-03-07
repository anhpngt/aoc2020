package common

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mocked returned variables for testPuzzle
var (
	errLoad     error
	part1Answer Answer
	part1Error  error
	part2Answer Answer
	part2Error  error
)

type testPuzzle struct{}

func (p *testPuzzle) Day() int {
	return 0
}

func (p *testPuzzle) Load(_ context.Context, datastream <-chan *LineContent) error {
	// Consume all the data from stream first
	for range datastream {
		continue
	}
	return errLoad
}

func (p *testPuzzle) SolvePart1() (Answer, error) {
	return part1Answer, part1Error
}

func (p *testPuzzle) SolvePart2() (Answer, error) {
	return part2Answer, part2Error
}

func TestSolveExample(t *testing.T) {
	errLoad = nil
	part1Answer = Answer("answer to part 1")
	part1Error = nil
	part2Answer = Answer("part 2's answer")
	part2Error = nil
	ans, err := SolveExample(&testPuzzle{})
	assert.NoError(t, err)
	assert.Equal(t, Answer("answer to part 1"), ans.First)
	assert.Equal(t, Answer("part 2's answer"), ans.Second)
}

func TestSolveExample_LoadError(t *testing.T) {
	errLoad = errors.New("some-loading-error")
	ans, err := SolveExample(&testPuzzle{})
	assert.EqualError(t, err, "failed to load input: some-loading-error")
	assert.Nil(t, ans)
}

func TestSolve(t *testing.T) {
	errLoad = nil
	part1Answer = Answer("answer to part 1")
	part1Error = nil
	part2Answer = Answer("part 2's answer")
	part2Error = nil
	ans, err := Solve(&testPuzzle{})
	assert.NoError(t, err)
	assert.Equal(t, Answer("answer to part 1"), ans.First)
	assert.Equal(t, Answer("part 2's answer"), ans.Second)
}

func TestSolve_LoadError(t *testing.T) {
	errLoad = errors.New("some-loading-error")
	ans, err := Solve(&testPuzzle{})
	assert.EqualError(t, err, "failed to load input: some-loading-error")
	assert.Nil(t, ans)
}

func TestSolve_UnableToSolvePart1(t *testing.T) {
	errLoad = nil
	part1Answer = ""
	part1Error = errors.New("part 1 is hard")
	ans, err := Solve(&testPuzzle{})
	assert.EqualError(t, err, "failed to solve part 1: part 1 is hard")
	assert.Nil(t, ans)
}

func TestSolve_UnableToSolvePart2(t *testing.T) {
	errLoad = nil
	part1Answer = "something"
	part1Error = nil
	part2Answer = ""
	part2Error = errors.New("need google's help")
	ans, err := Solve(&testPuzzle{})
	assert.EqualError(t, err, "failed to solve part 2: need google's help")
	assert.Nil(t, ans)
}
