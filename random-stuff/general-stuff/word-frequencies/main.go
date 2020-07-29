package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	msg = `
	***             Word Frequency calculator           ***
	usage: %s <file1> [<file2> [...<fileN>]]
	*******************************************************

	`
)

func main() {
	args := os.Args
	if len(args) == 1 || args[1] == "-h" || args[1] == "--help" {
		fmt.Fprintf(os.Stdout, msg+"\n", filepath.Base(args[0]))
	}

	frequencyPerWord := map[string]int{}
	for _, filename := range parseCmdLine(args[1:]) {
		updateFrequencies(filename, frequencyPerWord)
	}

	reportByWords(frequencyPerWord)
	wordsPerFrequency := invertMap(frequencyPerWord)
	reportByFrequency(wordsPerFrequency)
}

// parseCmdLine function predominantly handles the pattern match file globbing
// for windows platforms as the same is natural for POSIX systems
func parseCmdLine(files []string) []string {
	// to handle file globbing on Windows platform
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

// updateFrequencies function opens each file and hands it over to another
// function for the actual work/processing
func updateFrequencies(filename string, frequencyPerWord map[string]int) {
	var (
		file *os.File
		err  error
	)

	if file, err = os.Open(filename); err != nil {
		log.Printf("failed to open the file %v\n", err)
		return
	}

	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyPerWord)
}

// readAndUpdateFrequencies function takes each file reader and a map
// of frequencies per word
func readAndUpdateFrequencies(r *bufio.Reader, fw map[string]int) {
	for {
		line, err := r.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				fw[strings.ToLower(word)]++
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("failed to completely read the file %v\n", err)
			}
			break
		}
	}
}

// SplitOnNonLetters function splits a string at nonword characters
func SplitOnNonLetters(s string) []string {
	nonLetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FieldsFunc(s, nonLetter)
}

func reportByWords(fw map[string]int) {
	words := make([]string, 0, len(fw))
	var wordWidth, frequencyWidth int

	for word, frequency := range fw {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}

	// sort the slice
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, fw[word])
	}
}

func invertMap(intPerString map[string]int) map[int][]string {
	stringsPerInt := make(map[int][]string, len(intPerString))
	for key, value := range intPerString {
		stringsPerInt[value] = append(stringsPerInt[value], key)
	}
	return stringsPerInt
}

func reportByFrequency(wordsPerFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsPerFrequency))
	for f := range wordsPerFrequency {
		frequencies = append(frequencies, f)
	}

	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency -> Words")
	for _, f := range frequencies {
		words := wordsPerFrequency[f]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, f, strings.Join(words, ", "))
	}
}
