package main

/*
Count the occurrence of each word in a text file

- filename is passed as an argument to the program
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

// store the result of the words count into a map
var (
	results map[string]int
	line    string
	words   []string

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

	// printResults(results)
	showResults(results)
}

func wordProcess(words []string) map[string]int {
	for _, w := range words {
		w = re.ReplaceAllString(w, "")
		if w == "" {
			// eliminate the empty characters from the count
			continue
		}
		results[w] = results[w] + 1
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

func showResults(result map[string]int) {
	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Count", "Word")
	fmt.Fprintf(tw, format, "-----", "----")
	for w, c := range result {
		fmt.Fprintf(tw, format, c, w)
	}
	tw.Flush()
}
