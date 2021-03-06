package day2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	firstAns  common.Answer = "600"
	secondAns common.Answer = "245"
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

func TestSolutionDay1(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	require.Equal(t, firstAns, ans.First)
	require.Equal(t, secondAns, ans.Second)
}
