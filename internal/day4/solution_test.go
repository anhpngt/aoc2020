package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anhpngt/aoc2020/internal/common"
)

func TestPassport_MapKeyValue(t *testing.T) {
	p := &passport{}
	err := p.mapKeyValue("byr:2021")
	assert.NoError(t, err)

	err = p.mapKeyValue("byr2021")
	assert.EqualError(t, err, "invalid format for key:value pair: byr2021")

	assert.PanicsWithError(t, "invalid key in key:value pair: not_a_key", func() { _ = p.mapKeyValue("not_a_key:2021") })
}

func TestPassport_IsBYRValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{"1920", "1921", "1930", "1945", "1999", "2000", "2001", "2002"}
		for _, tc := range testcases {
			p := &passport{byr: tc}
			assert.Truef(t, p.isBYRValid(), "%s should be a valid birth year", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{"", "0", "0000", "1919", "2003", "abcd", "01999"}
		for _, tc := range testcases {
			p := &passport{byr: tc}
			assert.Falsef(t, p.isBYRValid(), "%s should not be a valid birth year", tc)
		}
	})
}

func TestPassport_IsIYRValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{"2010", "2011", "2019", "2020"}
		for _, tc := range testcases {
			p := &passport{iyr: tc}
			assert.Truef(t, p.isIYRValid(), "%s should be a valid issue year", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{"", "0", "0000", "2009", "2021", "abcd", "02010"}
		for _, tc := range testcases {
			p := &passport{iyr: tc}
			assert.Falsef(t, p.isIYRValid(), "%s should not be a valid issue year", tc)
		}
	})
}

func TestPassport_IsEYRValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{"2020", "2021", "2029", "2030"}
		for _, tc := range testcases {
			p := &passport{eyr: tc}
			assert.Truef(t, p.isEYRValid(), "%s should be a valid expiration year", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{"", "0", "0000", "2019", "2031", "abcd", "02020"}
		for _, tc := range testcases {
			p := &passport{eyr: tc}
			assert.Falsef(t, p.isEYRValid(), "%s should not be a valid expiration year", tc)
		}
	})
}

func TestPassport_IsHGTValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{
			"150cm", "151cm", "192cm", "193cm",
			"59in", "60in", "75in", "76in",
		}
		for _, tc := range testcases {
			p := &passport{hgt: tc}
			assert.Truef(t, p.isHGTValid(), "%s should be a valid height", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{
			"", "0", "0000",
			"59cm", "149cm", "194cm", "150cms",
			"58in", "77in", "150in", "59ins",
			"0150cm", "059in",
		}
		for _, tc := range testcases {
			p := &passport{hgt: tc}
			assert.Falsef(t, p.isHGTValid(), "%s should not be a valid height", tc)
		}
	})
}

func TestPassport_IsHCLValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{"#123abc", "#012789", "#abcdef"}
		for _, tc := range testcases {
			p := &passport{hcl: tc}
			assert.Truef(t, p.isHCLValid(), "%s should be a valid hair color", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{"", "#123abz", "123abc"}
		for _, tc := range testcases {
			p := &passport{hcl: tc}
			assert.Falsef(t, p.isHCLValid(), "%s should not be a valid hair color", tc)
		}
	})
}

func TestPassport_IsECLValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, tc := range testcases {
			p := &passport{ecl: tc}
			assert.Truef(t, p.isECLValid(), "%s should be a valid eye color", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{"", "abc", "xyz", "123", "wat"}
		for _, tc := range testcases {
			p := &passport{ecl: tc}
			assert.Falsef(t, p.isECLValid(), "%s should not be a valid eye color", tc)
		}
	})
}

func TestPassport_IsPIDValid(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testcases := []string{"000000000", "000000001", "999999999", "012345678", "123456789"}
		for _, tc := range testcases {
			p := &passport{pid: tc}
			assert.Truef(t, p.isPIDValid(), "%s should be a valid passport ID", tc)
		}
	})

	t.Run("negative", func(t *testing.T) {
		testcases := []string{"", "0000000000", "12345678", "0123456789"}
		for _, tc := range testcases {
			p := &passport{pid: tc}
			assert.Falsef(t, p.isPIDValid(), "%s should not be a valid passport ID", tc)
		}
	})
}

func TestSolveExample(t *testing.T) {
	ans, err := common.SolveExample(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("10"), ans.First)
	assert.Equal(t, common.Answer("6"), ans.Second)
}

func TestSolve(t *testing.T) {
	ans, err := common.Solve(&Puzzle{})
	require.NoError(t, err)
	assert.Equal(t, common.Answer("247"), ans.First)
	assert.Equal(t, common.Answer("145"), ans.Second)
}
