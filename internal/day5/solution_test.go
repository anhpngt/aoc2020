package day5

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestSeatCode_Decode(t *testing.T) {
	sc := seatCode("FBFBBFFRLR")
	sp, err := sc.decode()
	require.NoError(t, err)
	assert.Equal(t, 44, sp.row)
	assert.Equal(t, 5, sp.col)
}

func TestSeatPosition_ID(t *testing.T) {
	sc := seatCode("FBFBBFFRLR")
	sp, err := sc.decode()
	require.NoError(t, err)
	require.Equal(t, 357, sp.ID())

	sc = seatCode("BFFFBBFRRR")
	sp, err = sc.decode()
	require.NoError(t, err)
	require.Equal(t, 567, sp.ID())

	sc = seatCode("FFFBBBFRRR")
	sp, err = sc.decode()
	require.NoError(t, err)
	require.Equal(t, 119, sp.ID())

	sc = seatCode("BBFFBBFRLL")
	sp, err = sc.decode()
	require.NoError(t, err)
	require.Equal(t, 820, sp.ID())
}

func TestPuzzle_Load(t *testing.T) {
	p := &Puzzle{}
	out := make(chan *common.LineContent, 1)
	defer close(out)

	out <- &common.LineContent{Content: nil, Err: errors.New("cannot open file: a reason")}
	err := p.Load(context.Background(), out)
	assert.EqualError(t, err, "cannot open file: a reason")

	out <- &common.LineContent{Content: []byte("FFFFFFFRRRR")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, "seat code \"FFFFFFFRRRR\": invalid code length: 11")

	out <- &common.LineContent{Content: []byte("FLFFFFFRRR")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, "invalid character 'L' in seat code at position 1")

	out <- &common.LineContent{Content: []byte("FFFFFFFRRF")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, "invalid character 'F' in seat code at position 9")
}

func TestPuzzle_SolvePart2(t *testing.T) {
	p := &Puzzle{}
	p.spList = append(p.spList, seatPosition{0, 0})
	ans, err := p.SolvePart2() // should return ID of position (0, 1)
	require.NoError(t, err)
	require.Equal(t, common.ToAnswer((&seatPosition{0, 1}).ID()), ans)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("947"), ans.First)
	assert.Equal(t, common.Answer("636"), ans.Second)
}
