package main

/*
Count the occurrence of each word in a text file

- filename is passed as an argument to the program

- Create a list of words that shoudlnt' be included in the result.
	Examples of noise words are: 'on', 'a', 'the', 'are', 'in', 'of', etc.

- Running

	go run main.go data.txt sherlock-holmes.txt
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// store the result of the words count into a map
var (
	// map for holding the word and its count
	results map[string]int

	// create a regular expression matcher to match
	// anything except alphabets to filter the characters
	// like . , ; : etc
	re = regexp.MustCompile(`[^a-zA-Z]+`)

	wg           sync.WaitGroup
	resultsMutex sync.Mutex
)

const (
	usage = `
	Please provide input file(s) as argument
	usage: ./program <input file1> <input file2> | go run main.go <input file1> <input file2>
	`

	replaceChars = "`" + `~!@#$%^&*()-_+=[{]}\|;:'",<.>/?`
)

func main() {
	fmt.Println("--- 02 Counting the occurrence of each word in a text file ---")
	fmt.Println()

	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Printf("%s\n", usage)
		return
	}

	// record start time
	start := time.Now()
	log.Infof("Processing %v files for input", len(query))
	for i := range query {
		_, err := os.Stat(query[i])
		if os.IsNotExist(err) {
			log.Errorf("> input data file %q does not exist\n", query[i])
			return
		}
	}

	// initilize the map
	results = make(map[string]int)

	// process each input file
	for i := range query {
		processFiles(query[i])
	}

	wg.Wait()
	elapsed := time.Since(start)
	printResults(results)
	fmt.Printf("Total time: %v\n", elapsed)
}

// processFile function will process the contents of each input
// file one by and accumulates the word, count information into
// a soecified map
func processFiles(inputFile string) {
	wg.Add(1)

	// spawn an anonymous go routine
	go func() {
		log.Infof("Processing input file %v", inputFile)
		reader, err := os.Open(inputFile)
		if err != nil {
			log.Errorf("> Err: %q\n", err)
			return
		}

		in := bufio.NewScanner(reader)
		for in.Scan() {
			wordProcess(in.Text(), results)
		}

		if err := in.Err(); err != nil {
			log.Errorf("> Err: %q\n", err)
			return
		}

		wg.Done()
	}()
}

// wordProcess function processes each line from the text file
// and splits the words and filterd out the same of any unwanted
// chanarcters using regexp before pushing the word and it's
// associated count into the map
func wordProcess(input string, resmap map[string]int) {

	// for _, rep := range replaceChars {
	// 	input = strings.Replace(input, string(rep), "", -1)
	// }
	// input = strings.ToLower(input)
	// words := strings.Split(input, " ")

	line := strings.ToLower(input)
	words := strings.Fields(line)
	for _, w := range words {
		w = re.ReplaceAllString(w, "")
		// synchronize the map access so that concurrent update
		// of map is prevented
		resultsMutex.Lock()
		resmap[w] = resmap[w] + 1
		resultsMutex.Unlock()
	}
}

// printResults is a helper utility for printing the contents of
// the map in a visually sooting way
func printResults(result map[string]int) {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")
	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
