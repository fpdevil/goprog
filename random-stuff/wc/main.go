package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	usage = `
usage go run %s [files [file2...]]
`
)

// words struct for storing a map of word to occurrences
type words struct {
	sync.Mutex
	found map[string]int
}

// newWords function returns the words struct instance
func newWords() *words {
	return &words{found: map[string]int{}}
}

// add function helps to add a word and its count into words struct
func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	// if word alredy exists, increment count else just add
	// with the count of it
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// checkWords function scans the supplied file and adds each
// word with a default count of 1 to map
func checkWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

func main() {
	fmt.Println("* Word Counter Program *")
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, usage, filepath.Base(args[0]))
		return
	}

	var wg sync.WaitGroup
	// get a new instance of the struct
	w := newWords()
	// iterate through each supplied files
	for _, f := range args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := checkWords(file, w); err != nil {
				fmt.Fprintf(os.Stderr, "error parsing file %s\n", err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Printf("words appearing more than once:\n")
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Fprintf(os.Stdout, "%s: %d\n", word, count)
		}
	}
	w.Unlock()
}
