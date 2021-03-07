package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestSolutionDay1(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	require.Equal(t, common.Answer("545379"), ans.First)
	require.Equal(t, common.Answer("257778836"), ans.Second)
}
