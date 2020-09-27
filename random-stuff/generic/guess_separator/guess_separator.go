package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	Formatter := new(log.JSONFormatter)
	log.SetFormatter(Formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s file\n", filepath.Base(os.Args[0]))
		return
	}

	separators := []string{"\t", "*", "|", "•"}

	// read till first 5 lines of the file
	linesRead, lines := readTillNLines(os.Args[1], 5)
	counts := createCounts(lines, separators, linesRead)
	separator := guessSeparator(counts, separators, linesRead)
	report(separator)
}

// readTillNLines function reads the first N lines of a file
func readTillNLines(filename string, maxLines int) (int, []string) {
	logger := log.WithFields(log.Fields{
		"readTillNLines": "guess_separator",
	})

	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		logger.Fatalf("failed to open the file: %v", err)
	}
	defer file.Close()

	lines := make([]string, maxLines)
	reader := bufio.NewReader(file)

	i := 0
	for ; i < maxLines; i++ {
		// read till the first newline character
		line, err := reader.ReadString('\n')
		if line != "" {
			lines[i] = line
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Fatalf("failed to finish reading the file: %v", err)
		}
	}
	return i, lines[:i]
}

// createCounts function populates a matrix to hold the counts
// of each separator for each line read.
// each row represents the number of occurrences of the separator
// for each line
// size of matrix = no. of separators X no. of lines
func createCounts(lines, separators []string, linesRead int) [][]int {
	counts := make([][]int, len(separators))
	for sepIdx := range separators {
		counts[sepIdx] = make([]int, linesRead)
		for lineIdx, line := range lines {
			counts[sepIdx][lineIdx] = strings.Count(line, separators[sepIdx])
		}
	}
	return counts
}

// guessSeparator function as the name indicates discovers the separators in
// the lines from files. The function finds the first []int in counts slices
// whose counts are all the same and nonzero
// It iterates over each `row` of counts (which is per separator), and
// initially assumes that all the row's counts are the same.
func guessSeparator(counts [][]int, separators []string, linesRead int) string {
	for sepIdx := range separators {
		flag := true
		count := counts[sepIdx][0]
		for lineIdx := 1; lineIdx < linesRead; lineIdx++ {
			if counts[sepIdx][lineIdx] != count {
				flag = false
				break
			}
		}
		if count > 0 && flag {
			return separators[sepIdx]
		}
	}
	return ""
}

func report(separator string) {
	switch separator {
	case "":
		fmt.Println("whitespace separated or not separated at all!")
	case "\t":
		fmt.Println("tab separated!")
	default:
		fmt.Println("%s separated\n", separator)
	}
}
