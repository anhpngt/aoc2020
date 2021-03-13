package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAnswer(t *testing.T) {
	i := 1234567
	assert.Equal(t, Answer("1234567"), ToAnswer(i))
}

func TestToAnswerInt64(t *testing.T) {
	var i int64 = 123456789123456789
	assert.Equal(t, Answer("123456789123456789"), ToAnswerInt64(i))
}

func TestAnswerOfDay_String(t *testing.T) {
	ans := AnswerOfDay{1, Answer("abc"), Answer("123")}
	assert.Equal(t, "Day 1: abc and 123", ans.String())
}
