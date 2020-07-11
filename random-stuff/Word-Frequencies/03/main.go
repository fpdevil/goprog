package main

/*
Count the occurrence of each word in a text file

- filename is passed as an argument to the program

- Create a list of words that shoudlnt' be included in the result.
	Examples of noise words are: 'on', 'a', 'the', 'are', 'in', 'of', etc.
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// store the result of the words count into a map
var (
	results map[string]int
	line    string
	words   []string

	noisy = []string{"is", "to", "be", "in", "of", "for", "or", "not", "an", "are", "by"}

	// create a regular expression matcher to match
	// anything except alphabets to filter the characters
	// like . , ; : etc
	re = regexp.MustCompile(`[^a-zA-Z]+`)
)

const (
	usage = `
	Please provide input file as argument
	usage: ./program <input file> | go run main.go <input file>
	`
)

func main() {
	fmt.Println("--- 02 Counting the occurrence of each word in a text file ---")
	fmt.Println()

	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Printf("%s\n", usage)
		return
	}

	_, err := os.Stat(query[0])
	if os.IsNotExist(err) {
		fmt.Printf("input data file %s does not exist\n", query[0])
		return
	}

	reader, err := os.Open(query[0])
	if err != nil {
		fmt.Printf("> Err: %q\n", err)
		return
	}

	// initialize the map
	results = make(map[string]int)

	in := bufio.NewScanner(reader)
	for in.Scan() {
		line = strings.TrimSpace(in.Text())
		words = strings.Fields(line)
		wordProcess(words)
	}

	// error handling for the scanner
	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %q\n", err)
		return
	}

	printResults(results)
}

func wordProcess(words []string) map[string]int {
	var match bool
	for _, w := range words {
		w = re.ReplaceAllString(w, "")
		for _, f := range noisy {
			if strings.ToLower(w) == strings.ToLower(f) {
				match = true
				break
			}
		}

		if !match {
			results[w] = results[w] + 1
		}
		match = false
	}
	return results
}

func printResults(result map[string]int) {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")
	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
