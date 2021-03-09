package day6

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestMCQ_AddFrom(t *testing.T) {
	assert := assert.New(t)
	m := &mcq{}
	err := m.addFrom("!")
	assert.EqualError(err, "invalid character \"!\" in yes-question string")
	assert.Equal(0, m.groupsize)
	for _, v := range m.answer {
		assert.Equal(0, v)
	}
}

func TestPuzzle_Load(t *testing.T) {
	p := &Puzzle{}
	out := make(chan *common.LineContent, 1)
	defer close(out)

	out <- &common.LineContent{Content: nil, Err: errors.New("cannot open file: a reason")}
	err := p.Load(context.Background(), out)
	assert.EqualError(t, err, "cannot open file: a reason")

	out <- &common.LineContent{Content: []byte("")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, "unexpected blank line in input file")
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("11"), ans.First)
	assert.Equal(t, common.Answer("6"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("6763"), ans.First)
	assert.Equal(t, common.Answer("3512"), ans.Second)
}
