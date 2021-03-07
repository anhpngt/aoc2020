package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestValidatePosition(t *testing.T) {
	testcase := []struct {
		pw     password
		pwp    passwordPolicy
		expect bool
	}{
		{
			pw:     []rune("abcde"),
			pwp:    passwordPolicy{1, 3, 'a'},
			expect: true,
		},
		{
			pw:     []rune("cdefg"),
			pwp:    passwordPolicy{1, 3, 'b'},
			expect: false,
		},
		{
			pw:     []rune("ccccccccc"),
			pwp:    passwordPolicy{2, 9, 'c'},
			expect: false,
		},
	}

	for _, tc := range testcase {
		require.Equal(t, tc.expect, tc.pwp.validatePosition(tc.pw))
	}
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("2"), ans.First)
	assert.Equal(t, common.Answer("1"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("600"), ans.First)
	assert.Equal(t, common.Answer("245"), ans.Second)
}
