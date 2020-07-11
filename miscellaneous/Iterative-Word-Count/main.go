package main

/*
Write a function  to count the frequency of words  in a string of
text and  return a map of  words with their counts.  The function
should convert the  text to lowercase, and  punctuation should be
trimmed from words.
*/

import (
	"fmt"
	"regexp"
	"strings"
)

const text = `
As far as eye could reach he saw nothing but the stems of the
great plants about him receding in the violet shade, and far overhead
the multiple transparency of huge leaves filtering the sunshine to the
solemn splendour of twilight in which he walked. Whenever he felt able
he ran again; the ground continued soft and springy, covered with the
same resilient weed which was the first thing his hands had touched in
Malacandra. Once or twice a small red creature scuttled across his
path, but otherwise there seemed to be no life stirring in the wood;
nothing to fear —- except the fact of wandering unprovisioned and alone
in a forest of unknown vegetation thousands or millions of miles
beyond the reach or knowledge of man.
`

// countWords function will get a count of each word in the text
func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))

	// create a regular expression matcher to match
	// anything except alphabets to filter the characters
	// like . , ; : etc
	re := regexp.MustCompile(`[^a-z]+`)

	// define a map for holding the return data
	frequency := make(map[string]int, len(words))

	for _, word := range words {
		word := re.ReplaceAllString(word, "")
		frequency[word]++
	}
	return frequency
}

func main() {
	fmt.Println("--- frequency of words in a passage ---")
	frequency := countWords(text)

	for k, v := range frequency {
		if v > 1 {
			fmt.Printf("%v: %d\n", k, v)
		}
	}
}
