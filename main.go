package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/anhpngt/aoc2020/internal/common"
	"github.com/anhpngt/aoc2020/internal/day1"
	"github.com/anhpngt/aoc2020/internal/day2"
	"github.com/anhpngt/aoc2020/internal/day3"
	"github.com/anhpngt/aoc2020/internal/day4"
	"github.com/anhpngt/aoc2020/internal/day5"
	"github.com/anhpngt/aoc2020/internal/day6"
	"github.com/anhpngt/aoc2020/internal/day7"
)

// NumDaySolved is the number of days this solution has solved.
const NumDaySolved = 7

var (
	help      bool
	targetDay int
)

func init() {
	flag.BoolVar(&help, "h", false, "print help message")
	flag.IntVar(&targetDay, "d", 0, fmt.Sprintf("day number to solve, must be in the range of 0 and %d inclusive", NumDaySolved))
}

func main() {
	processFlag()

	if !(0 < targetDay && targetDay <= NumDaySolved) {
		printUsage()
		panic(fmt.Errorf("invalid target day to solve, must be in the range of 0 and %d inclusive", NumDaySolved))
	}

	solve(targetDay)
}

func processFlag() {
	flag.Parse()

	if help {
		printUsage()
		os.Exit(0)
	}

	if flag.NArg() > 0 {
		var err error
		targetDay, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			printUsage()
			os.Exit(1)
		}
	}
}

func printUsage() {
	fmt.Println("Usage: go run main.go -d TARGET_DAY")
	flag.PrintDefaults()
}

func solve(n int) {
	fmt.Printf("Solving day %d's puzzle\n", n)
	var p common.Puzzle
	switch n {
	case 1:
		p = &day1.Puzzle{}
	case 2:
		p = &day2.Puzzle{}
	case 3:
		p = &day3.Puzzle{}
	case 4:
		p = &day4.Puzzle{}
	case 5:
		p = &day5.Puzzle{}
	case 6:
		p = &day6.Puzzle{}
	case 7:
		p = &day7.Puzzle{}
	}

	ans, err := common.Solve(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}
