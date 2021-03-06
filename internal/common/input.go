package common

import (
	"bufio"
	"context"
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

// getInputFilename returns absolute path to the puzzle input of day n.
func getInputFilename(n int) string {
	return filepath.Join(inputDirectory, fmt.Sprintf("%d.txt", n))
}

// LineContent contains data of a line in the input file. It also contains the error
// for inter-channel communication.
type LineContent struct {
	Content []byte
	Err     error
}

// LoadInputAsync ...
func LoadInputAsync(ctx context.Context, day, chanSize int) <-chan LineContent {
	if chanSize < 0 {
		chanSize = 0
	}
	out := make(chan LineContent, chanSize)
	go func() {
		defer close(out)

		filename := getInputFilename(day)
		file, err := os.Open(filename)
		if err != nil {
			select {
			case <-ctx.Done():
				return
			case out <- LineContent{nil, err}:
			}
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case out <- LineContent{scanner.Bytes(), nil}:
			}
		}

		if err := scanner.Err(); err != nil {
			select {
			case <-ctx.Done():
				return
			case out <- LineContent{nil, fmt.Errorf("an error occurred while reading from input file: %s", err)}:
			}
		}
	}()
	return out
}
