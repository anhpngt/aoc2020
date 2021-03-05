package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/anhpngt/aoc2020/internal/common"
	"github.com/anhpngt/aoc2020/internal/day1"
)

// NumDaySolved is the number of days this solution has solved.
const NumDaySolved = 1

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
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("Usage: go run main.go -d [TARGET]")
}

func solve(n int) {
	fmt.Printf("Solving day %d\n", n)
	var p common.Puzzle
	switch n {
	case 1:
		p = &day1.Puzzle{}
	}

	ans, err := common.Solve(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}
