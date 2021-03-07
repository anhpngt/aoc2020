package common

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// fileScanBufferSize is the buffer size while scanning input files, in KB.
const fileScanBufferSize = 1024

var inputDirectory string

func init() {
	_, fn, _, ok := runtime.Caller(0)
	if !ok {
		panic("common.init() failed: cannot get config file path")
	}
	inputDirectory = filepath.Join(filepath.Dir(fn), "../../input")
}

// LineContent contains data of a line in the input file. It also contains the error
// for inter-channel communication.
type LineContent struct {
	Content []byte
	Err     error
}

// getInputFilename returns absolute path to the puzzle input of day n.
func getInputFilename(n int) string {
	return filepath.Join(inputDirectory, fmt.Sprintf("%d.txt", n))
}

// getExampleInputFilename returns absolute path to the puzzle example input of day n.
func getExampleInputFilename(n int) string {
	return filepath.Join(inputDirectory, fmt.Sprintf("%d.example.txt", n))
}

// loadInputAsync reads puzzle input for the day and passes each line from the file to the
// returned channel. The channel is closed by the function itself when the reading is finished,
// or when the context is canceled.
func loadInputAsync(ctx context.Context, day, chanSize int) <-chan *LineContent {
	filename := getInputFilename(day)
	return loadFileAsync(ctx, filename, chanSize)
}

// loadExampleInputAsync reads puzzle example input for the day and passes each line from the
// file to the returned channel. The channel is closed by the function itself when the reading
// is finished, or when the context is canceled.
func loadExampleInputAsync(ctx context.Context, day, chanSize int) <-chan *LineContent {
	filename := getExampleInputFilename(day)
	return loadFileAsync(ctx, filename, chanSize)
}

// loadFileAsync reads from file and passes each line from the file to the returned channel.
// The channel is closed by the function itself when the reading is finished, or when the context
// is canceled.
func loadFileAsync(ctx context.Context, filename string, chanSize int) <-chan *LineContent {
	if chanSize < 0 {
		chanSize = 0
	}
	out := make(chan *LineContent, chanSize)
	go func() {
		defer close(out)

		file, err := os.Open(filename)
		if err != nil {
			select {
			case <-ctx.Done():
			case out <- &LineContent{nil, err}:
			}
			return
		}

		scanner := bufio.NewScanner(file)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, fileScanBufferSize*1024) // increase the buffer size from 64 KB
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case out <- &LineContent{scanner.Bytes(), nil}:
			}
		}

		if err := scanner.Err(); err != nil {
			select {
			case <-ctx.Done():
				return
			case out <- &LineContent{nil, fmt.Errorf("an error occurred while reading from input file: %s", err)}:
			}
		}
	}()
	return out
}
