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
)

// NumDaySolved is the number of days this solution has solved.
const NumDaySolved = 3

var (
	help      bool
	targetDay int
)

func init() {
	flag.BoolVar(&help, "h", false, "print help message")
	flag.IntVar(&targetDay, "d", 0, fmt.Sprintf("day number to solve, must be greater than 0 and less than %d", NumDaySolved+1))
}

func main() {
	processFlag()

	if !(0 < targetDay && targetDay <= NumDaySolved) {
		printUsage()
		panic(fmt.Errorf("invalid target day to solve, must be greater than 0 and less than %d", NumDaySolved+1))
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
	}

	ans, err := common.Solve(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}
