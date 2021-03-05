package common

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var inputDirectory string

func init() {
	_, fn, _, ok := runtime.Caller(0)
	if !ok {
		panic("common.init() failed: cannot get config file path")
	}
	inputDirectory = filepath.Join(filepath.Dir(fn), "../../input")
}

// LoadPuzzleInput reads puzzle input of day n's problem.
func LoadPuzzleInput(n int) ([]byte, error) {
	filename := fmt.Sprintf("%d.txt", n)
	filepathabs := filepath.Join(inputDirectory, filename)
	return os.ReadFile(filepathabs)
}
