package day1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestPuzzle_Load(t *testing.T) {
	p := &Puzzle{}
	out := make(chan *common.LineContent, 1)
	defer close(out)

	out <- &common.LineContent{Content: nil, Err: errors.New("cannot open file: a reason")}
	err := p.Load(context.Background(), out)
	assert.EqualError(t, err, "cannot open file: a reason")

	out <- &common.LineContent{Content: []byte("not a number"), Err: nil}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, "invalid line input: strconv.Atoi: parsing \"not a number\": invalid syntax")
}

func TestPuzzle_SolvePart1(t *testing.T) {
	t.Run("impossible case", func(t *testing.T) {
		p := &Puzzle{expenses: []int{1, 2, 3}}
		ans, err := p.SolvePart1()
		assert.EqualError(t, err, common.ErrCannotComputeAnswer.Error())
		assert.Empty(t, ans)
	})
}

func TestPuzzle_SolvePart2(t *testing.T) {
	t.Run("impossible case", func(t *testing.T) {
		p := &Puzzle{expenses: []int{1, 2, 3}}
		ans, err := p.SolvePart2()
		assert.EqualError(t, err, common.ErrCannotComputeAnswer.Error())
		assert.Empty(t, ans)
	})
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("514579"), ans.First)
	assert.Equal(t, common.Answer("241861950"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("545379"), ans.First)
	assert.Equal(t, common.Answer("257778836"), ans.Second)
}
