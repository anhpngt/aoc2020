package day11

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
}

func TestPuzzle_LoadEmpty(t *testing.T) {
	p := &Puzzle{}
	out := make(chan *common.LineContent, 1)
	close(out)
	err := p.Load(context.Background(), out)
	assert.EqualError(t, err, "empty input data")
}

func TestPuzzle_LoadInvalidSize(t *testing.T) {
	p := &Puzzle{}
	out := make(chan *common.LineContent, 2)
	out <- &common.LineContent{Content: []byte("LL"), Err: nil}
	out <- &common.LineContent{Content: []byte("LLL"), Err: nil}
	close(out)
	err := p.Load(context.Background(), out)
	assert.EqualError(t, err, "column length mismatched for row \"LLL\": expected 2, got 3")
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("37"), ans.First)
	assert.Equal(t, common.Answer("26"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("2424"), ans.First)
	assert.Equal(t, common.Answer("2208"), ans.Second)
}
