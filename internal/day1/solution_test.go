package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	firstAns  common.Answer = "545379"
	secondAns common.Answer = "257778836"
)

func TestSolutionDay1(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	require.Equal(t, firstAns, ans.First)
	require.Equal(t, secondAns, ans.Second)
}
