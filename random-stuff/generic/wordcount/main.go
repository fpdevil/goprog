package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Words represents each word and its cont
type Words struct {
	sync.Mutex
	found map[string]int
}

// New initializes and cerates an in instance of words
func New() *Words {
	return &Words{found: map[string]int{}}
}

// Add function populates the Word struct with words if they are not
// present and gets the count if it already contains one
func (w *Words) Add(word string, n int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// checkWords function parses the input file contents and counts the
// words which appear more than once in it
func checkWords(filename string, dict *Words) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		word := strings.ToLower(sc.Text())
		dict.Add(word, 1)
	}
	return sc.Err()
}

func main() {
	var wg sync.WaitGroup

	w := New()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := checkWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}

	wg.Wait()
	fmt.Printf("Words appearing more than once\n\n")

	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()
}
