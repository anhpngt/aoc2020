package day7

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
)

const (
	dayNumber = 7

	targetColor = "shiny gold"
)

type (
	nodeStatus  int
	countStatus int
)

const (
	statusUnknown nodeStatus = iota
	statusFound
	statusNotFound

	statusUncounted countStatus = iota
	statusCounted
)

var (
	ruleRe         = regexp.MustCompile(`^([a-z ]+) bags contain ([a-z0-9 ,]+)\.$`)
	objectPhraseRe = regexp.MustCompile(`(?U)(\d+) ([a-z ]+) bags?`)

	emptyObjectPhrase = []byte("no other bags")
)

type node struct {
	name      string
	children  map[string]int
	hasTarget nodeStatus

	countStatus countStatus
	count       int
}

func newNode(b []byte) (*node, error) {
	nd := &node{
		children:    make(map[string]int),
		hasTarget:   statusUnknown,
		countStatus: statusUncounted,
	}
	err := nd.parseRule(b)
	if err != nil {
		return nil, err
	}
	return nd, nil
}

// parseRule parses one complete rule into a node.
func (nd *node) parseRule(r []byte) error {
	matches := ruleRe.FindAllSubmatch(r, -1)
	if matches == nil {
		return errors.New(`rule format "<color> bags contain <objects>" violated`)
	}

	nd.name = string(matches[0][1])
	err := nd.parseObjectPhrase(matches[0][2])
	return err
}

// parseContainPhrase parses the target-section of the rule, e.g. "no other bags", or
// "1 bright white bag, 2 muted yellow bags", etc..., into node's children.
func (nd *node) parseObjectPhrase(p []byte) error {
	if bytes.Compare(p, emptyObjectPhrase) == 0 {
		return nil
	}
	allmatches := objectPhraseRe.FindAllSubmatch(p, -1)
	if allmatches == nil {
		return errors.New("cannot find valid bags in object phrase")
	}

	for _, match := range allmatches {
		quantity, _ := strconv.Atoi(string(match[1]))
		if quantity == 0 {
			return errors.New("bag quantity cannot not be zero")
		}
		color := string(match[2])
		if _, existed := nd.children[color]; existed {
			return fmt.Errorf("found duplicate color in rule: %s", color)
		}
		nd.children[color] = quantity
	}
	return nil
}

func (nd *node) searchForTarget(graph map[string]*node) bool {
	if nd.hasTarget != statusUnknown {
		return nd.hasTarget == statusFound
	}

	// DFS
	for childName := range nd.children {
		child := graph[childName]
		if child.searchForTarget(graph) {
			nd.hasTarget = statusFound
			return true
		}
	}

	nd.hasTarget = statusNotFound
	return false
}

func (nd *node) countContainingBags(graph map[string]*node) int {
	if nd.countStatus == statusCounted {
		return nd.count + 1
	}

	for childName, childQuantity := range nd.children {
		child := graph[childName]
		childCount := child.countContainingBags(graph)
		nd.count += childCount * childQuantity
	}

	nd.countStatus = statusCounted
	return nd.count + 1 // plus 1 for the node's bag it self
}

// Puzzle contains the puzzle for day 7.
type Puzzle struct {
	lookup map[string]*node
}

// Day returns 7.
func (p *Puzzle) Day() int {
	return dayNumber
}

// Load loads the puzzle input for day 7.
func (p *Puzzle) Load(ctx context.Context, datastream <-chan *common.LineContent) error {
	p.lookup = make(map[string]*node)
	for dataline := range datastream {
		if dataline.Err != nil {
			return dataline.Err
		}

		nd, err := newNode(dataline.Content)
		if err != nil {
			return fmt.Errorf("failed to parse rule \"%s\": %s", string(dataline.Content), err)
		}
		if _, existed := p.lookup[nd.name]; existed {
			return fmt.Errorf("found duplicate rules for \"%s\"", nd.name)
		}
		p.lookup[nd.name] = nd
	}

	return nil
}

// SolvePart1 returns the answer to day 7, part 1.
func (p *Puzzle) SolvePart1() (common.Answer, error) {
	count := 0
	target := p.lookup[targetColor]
	target.hasTarget = statusFound

	// Perform a search on the target
	for color := range p.lookup {
		if color == targetColor {
			continue
		}

		nd := p.lookup[color]
		if nd.searchForTarget(p.lookup) {
			count++
		}
	}
	return common.ToAnswer(count), nil
}

// SolvePart2 returns the answer to day 7, part 2.
func (p *Puzzle) SolvePart2() (common.Answer, error) {
	target := p.lookup[targetColor]
	count := target.countContainingBags(p.lookup) - 1 // minus 1 from the target bag itself
	return common.ToAnswer(count), nil
}
