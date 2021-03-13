package day9

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
	assert.EqualError(t, err, `invalid data line: strconv.Atoi: parsing "not a number": invalid syntax`)
}

func TestSolveExample(t *testing.T) {
	p := &Puzzle{blockSize: 5}
	ans, err := common.SolveExample(p)
	require.NoError(t, err)
	assert.Equal(t, common.Answer("127"), ans.First)
	assert.Equal(t, common.Answer("62"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("15690279"), ans.First)
	assert.Equal(t, common.Answer("2174232"), ans.Second)
}

func TestPuzzle_SolvePart1_Failure(t *testing.T) {
	p := &Puzzle{
		blockSize: 2,
		data:      []int{0, 1, 1},
	}
	ans, err := p.SolvePart1()
	assert.EqualError(t, err, "could not find the answer")
	assert.Equal(t, common.Answer(""), ans)
}

func TestPuzzle_SolvePart2_SetBlockSize(t *testing.T) {
	p := &Puzzle{
		blockSize:     0,
		data:          []int{0, 0, 0},
		invalidNumber: 1,
	}
	ans, err := p.SolvePart2()
	assert.EqualError(t, err, "could not find the answer")
	assert.Equal(t, common.Answer(""), ans)
}
