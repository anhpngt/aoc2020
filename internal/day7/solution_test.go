package day7

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

	out <- &common.LineContent{Content: []byte("some invalid rule")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, `failed to parse rule "some invalid rule": rule format "<color> bags contain <objects>" violated`)

	out <- &common.LineContent{Content: []byte("red bags contain blue bag.")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, `failed to parse rule "red bags contain blue bag.": cannot find valid bags in object phrase`)

	out <- &common.LineContent{Content: []byte("red bags contain 0 blue bags.")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, `failed to parse rule "red bags contain 0 blue bags.": bag quantity cannot not be zero`)

	out <- &common.LineContent{Content: []byte("red bags contain 1 blue bag, 2 blue bags.")}
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, `failed to parse rule "red bags contain 1 blue bag, 2 blue bags.": found duplicate color in rule: blue`)

	go func() {
		out <- &common.LineContent{Content: []byte("red bags contain 1 blue bag.")}
		out <- &common.LineContent{Content: []byte("red bags contain 2 green bags.")}
	}()
	err = p.Load(context.Background(), out)
	assert.EqualError(t, err, `found duplicate rules for "red"`)
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("4"), ans.First)
	assert.Equal(t, common.Answer("32"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("211"), ans.First)
	assert.Equal(t, common.Answer("12414"), ans.Second)
}
