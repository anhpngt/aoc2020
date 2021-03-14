package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeading_ToValue(t *testing.T) {
	testcases := []struct {
		in  rune
		out int
	}{
		{'N', 270},
		{'W', 180},
		{'S', 90},
		{'E', 0},
	}

	assert := assert.New(t)
	for _, tc := range testcases {
		h := heading(tc.in)
		assert.Equal(tc.out, h.toValue())
	}
}

func TestHeading_ToValue_InvalidHeading(t *testing.T) {
	h := heading('x')
	assert.PanicsWithError(t,
		"invalid heading value: x",
		func() { _ = h.toValue() },
	)
}

func TestHeading_FromValue(t *testing.T) {
	testcases := []struct {
		in  int
		out rune
	}{
		{540, 'W'},
		{270, 'N'},
		{180, 'W'},
		{90, 'S'},
		{0, 'E'},
		{-90, 'N'},
	}

	assert := assert.New(t)
	for _, tc := range testcases {
		h := heading('E')
		h.fromValue(tc.in)
		assert.Equal(tc.out, rune(h))
	}
}

func TestHeading_FromValue_InvalidValue(t *testing.T) {
	h := heading('E')
	assert.PanicsWithError(t,
		"invalid value for heading: 1",
		func() { h.fromValue(1) },
	)
}

func TestPosition_Move(t *testing.T) {
	p := initialPosition()
	assert.PanicsWithError(t,
		"invalid action: A",
		func() { p.move('A', 0) },
	)
}

func TestPositionWithWaypoint_Move(t *testing.T) {
	p := initialTestPositionWithWaypoint()
	assert.PanicsWithError(t,
		"invalid action: A",
		func() { p.move('A', 0) },
	)
}

func TestPositionWithWaypoint_TurnWaypointLeft(t *testing.T) {
	testcases := []struct {
		in  int
		out position
	}{
		{-90, position{1, -10, heading('E')}},
		{0, position{10, 1, heading('E')}},
		{90, position{-1, 10, heading('E')}},
		{180, position{-10, -1, heading('E')}},
		{270, position{1, -10, heading('E')}},
		{360, position{10, 1, heading('E')}},
		{450, position{-1, 10, heading('E')}},
	}

	assert := assert.New(t)
	for _, tc := range testcases {
		p := initialTestPositionWithWaypoint()
		p.turnWaypointLeft(tc.in)
		assert.Equalf(tc.out, p.wp, "wrong answer for input %d", tc.in)
	}
}

func TestPositionWithWaypoint_TurnWaypointLeft_InvalidValue(t *testing.T) {
	p := initialTestPositionWithWaypoint()
	assert.PanicsWithError(t,
		"invalid value for action: 1",
		func() { p.turnWaypointLeft(1) },
	)
}

func TestPositionWithWaypoint_TurnWaypointRight(t *testing.T) {
	testcases := []struct {
		in  int
		out position
	}{
		{-90, position{-1, 10, heading('E')}},
		{0, position{10, 1, heading('E')}},
		{90, position{1, -10, heading('E')}},
		{180, position{-10, -1, heading('E')}},
		{270, position{-1, 10, heading('E')}},
		{360, position{10, 1, heading('E')}},
		{450, position{1, -10, heading('E')}},
	}

	assert := assert.New(t)
	for _, tc := range testcases {
		p := initialTestPositionWithWaypoint()
		p.turnWaypointRight(tc.in)
		assert.Equalf(tc.out, p.wp, "wrong answer for input %d", tc.in)
	}
}

func TestPositionWithWaypoint_TurnWaypointRight_InvalidValue(t *testing.T) {
	p := initialTestPositionWithWaypoint()
	assert.PanicsWithError(t,
		"invalid value for action: 1",
		func() { p.turnWaypointRight(1) },
	)
}

func initialTestPositionWithWaypoint() positionWithWaypoint {
	return positionWithWaypoint{
		pos: initialPosition(),
		wp:  position{10, 1, heading('E')},
	}
}
