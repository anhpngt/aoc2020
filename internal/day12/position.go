package day12

import "fmt"

type heading rune

func (h *heading) toValue() int {
	switch *h {
	case 'N':
		return 270
	case 'W':
		return 180
	case 'S':
		return 90
	case 'E':
		return 0
	default:
		panic(fmt.Errorf("invalid heading value: %s", string(*h)))
	}
}

func (h *heading) fromValue(v int) {
	normalizeDegree(&v)
	switch v {
	case 270:
		*h = 'N'
	case 180:
		*h = 'W'
	case 90:
		*h = 'S'
	case 0:
		*h = 'E'
	default:
		panic(fmt.Errorf("invalid value for heading: %d", v))
	}
}

func (h *heading) turnLeft(v int) {
	newValue := h.toValue() - v
	h.fromValue(newValue)
}

func (h *heading) turnRight(v int) {
	newValue := h.toValue() + v
	h.fromValue(newValue)
}

type position struct {
	x   int
	y   int
	hdg heading
}

func (p *position) move(action rune, value int) {
	switch action {
	case 'N':
		p.y += value
	case 'S':
		p.y -= value
	case 'E':
		p.x += value
	case 'W':
		p.x -= value
	case 'L':
		p.hdg.turnLeft(value)
	case 'R':
		p.hdg.turnRight(value)
	case 'F':
		p.move(rune(p.hdg), value)
	default:
		panic(fmt.Errorf("invalid action: %s", string(action)))
	}
}

func (p *position) mhtDist() int {
	return abs(p.x) + abs(p.y)
}

type positionWithWaypoint struct {
	pos position
	wp  position
}

func (p *positionWithWaypoint) move(action rune, value int) {
	switch action {
	case 'N', 'S', 'E', 'W':
		p.wp.move(action, value)
	case 'L':
		p.turnWaypointLeft(value)
	case 'R':
		p.turnWaypointRight(value)
	case 'F':
		p.pos.x += p.wp.x * value
		p.pos.y += p.wp.y * value
	default:
		panic(fmt.Errorf("invalid action: %s", string(action)))
	}
}

func (p *positionWithWaypoint) turnWaypointLeft(v int) {
	normalizeDegree(&v)
	var a, b, c, d int
	switch v {
	case 0:
		a, b, c, d = 1, 0, 0, 1
	case 90:
		a, b, c, d = 0, -1, 1, 0
	case 180:
		a, b, c, d = -1, 0, 0, -1
	case 270:
		a, b, c, d = 0, 1, -1, 0
	default:
		panic(fmt.Errorf("invalid value for action: %d", v))
	}
	p.wp.x, p.wp.y = applyRotationMatrix(p.wp.x, p.wp.y, a, b, c, d)
}

func (p *positionWithWaypoint) turnWaypointRight(v int) {
	normalizeDegree(&v)
	var a, b, c, d int
	switch v {
	case 0:
		a, b, c, d = 1, 0, 0, 1
	case 90:
		a, b, c, d = 0, 1, -1, 0
	case 180:
		a, b, c, d = -1, 0, 0, -1
	case 270:
		a, b, c, d = 0, -1, 1, 0
	default:
		panic(fmt.Errorf("invalid value for action: %d", v))
	}
	p.wp.x, p.wp.y = applyRotationMatrix(p.wp.x, p.wp.y, a, b, c, d)
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func normalizeDegree(v *int) {
	for *v < 0 {
		*v += 360
	}
	*v %= 360
}

func applyRotationMatrix(x, y, a, b, c, d int) (int, int) {
	return a*x + b*y, c*x + d*y
}
