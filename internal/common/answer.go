package common

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrCannotComputeAnswer happens when the algorithm cannot find the answer.
var ErrCannotComputeAnswer = errors.New("cannot compute the answer")

// Answer contains the answer to a single question.
type Answer string

// ToAnswer converts an integer to Answer.
func ToAnswer(i int) Answer {
	return Answer(strconv.Itoa(i))
}

// AnswerOfDay contains the answer to a day's puzzle.
type AnswerOfDay struct {
	Day    int
	First  Answer
	Second Answer
}

func (a *AnswerOfDay) String() string {
	return fmt.Sprintf("Day %d: %s and %s", a.Day, a.First, a.Second)
}
