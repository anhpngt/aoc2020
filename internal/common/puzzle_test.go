package common

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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
	return nil
}

func (p *testPuzzle) SolvePart1() (Answer, error) {
	return Answer("answer to part 1"), nil
}

func (p *testPuzzle) SolvePart2() (Answer, error) {
	return Answer("part 2's answer"), nil
}

func TestSolveExample(t *testing.T) {
	ans, err := SolveExample(&testPuzzle{})
	assert.NoError(t, err)
	assert.Equal(t, Answer("answer to part 1"), ans.First)
	assert.Equal(t, Answer("part 2's answer"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := Solve(&testPuzzle{})
	assert.NoError(t, err)
	assert.Equal(t, Answer("answer to part 1"), ans.First)
	assert.Equal(t, Answer("part 2's answer"), ans.Second)
}
