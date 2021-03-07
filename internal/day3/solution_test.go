package day3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestNewGrid(t *testing.T) {
	testcases := []struct {
		nrows int
		ncols int
		data  [][]rune
		err   error
	}{
		{2, 2, [][]rune{[]rune("12"), []rune("34")}, nil},
		{3, 2, [][]rune{[]rune("12"), []rune("34")}, errors.New("grid data size mismatched: expecting 3 rows, got 2 rows")},
		{2, 2, [][]rune{[]rune("12")}, errors.New("grid data size mismatched: expecting 2 rows, got 1 rows")},
		{2, 2, [][]rune{[]rune("12"), []rune("345")}, errors.New(`grid data size mismatched: expecting 2 columns, got 3 columns for "345"`)},
	}

	for _, tc := range testcases {
		resp, err := newGrid(tc.nrows, tc.ncols, tc.data)
		if tc.err == nil {
			require.NoError(t, err)
			require.NotNil(t, resp)
		} else {
			require.EqualError(t, err, tc.err.Error())
			require.Nil(t, resp)
		}
	}
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("7"), ans.First)
	assert.Equal(t, common.Answer("336"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("259"), ans.First)
	assert.Equal(t, common.Answer("2224913600"), ans.Second)
}
