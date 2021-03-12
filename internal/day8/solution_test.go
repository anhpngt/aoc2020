package day8

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

	out <- &common.LineContent{Content: []byte("dummy instruction"), Err: nil}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, `failed to parse instruction "dummy instruction": invalid format`)
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("5"), ans.First)
	assert.Equal(t, common.Answer("8"), ans.Second)
}

func TestSolveExample_2(t *testing.T) {
	p := &Puzzle{
		instructions: []*op{
			{nop, 2},
			{jmp, 0},
		},
	}
	ans, err := p.SolvePart2()
	require.NoError(t, err)
	assert.Equal(t, common.Answer("0"), ans)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("1087"), ans.First)
	assert.Equal(t, common.Answer("780"), ans.Second)
}
