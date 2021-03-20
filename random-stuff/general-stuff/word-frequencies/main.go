package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
)

const (
	msg = `
	***             Word Frequency calculator           ***
	usage: %s <file1> [<file2> [...<fileN>]]
	*******************************************************

	`
)

func init() {
	Formatter := new(log.JSONFormatter)
	log.SetFormatter(Formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	logger := log.WithFields(log.Fields{
		"main": "word-frequencies",
	})
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "help" {
		logger.Errorf("%s started without arguments", filepath.Base(os.Args[0]))
		fmt.Printf(msg, filepath.Base(os.Args[0]))
		return
	}
	logger.Info("starting Word Frequency service...")

	// define a map for holding the word and its corresponding frequency
	frequencyOfWord := map[string]int{}

	// parse the command line arguments to take the input files
	for _, filename := range parseCmdLine(os.Args) {
		updateFrequencies(filename, frequencyOfWord)
	}

	reportByWords(frequencyOfWord)
	wordsPerFrequency := invertMap(frequencyOfWord)
	reportByFrequency(wordsPerFrequency)
}

// parseCmdLine function predominantly handles the pattern match file globbing
// for windows platforms as the same is natural and default for POSIX systems
// The filepath.Glob returns the names of all files matching pattern or nil if
// there is no matching file.
// an example pattern to look for may be /usr/*/bin/sys
func parseCmdLine(args []string) []string {
	logger := log.WithFields(log.Fields{
		"parseCmdLine": "word-frequencies",
	})
	if runtime.GOOS == "windows" {
		files := make([]string, 0, len(args))
		for _, name := range args {
			if matches, err := filepath.Glob(name); err != nil {
				logger.Errorf("invalid matching path for the glob %s", err.Error())
				files = append(files, name)
			} else if matches != nil {
				logger.Debugf("found match for %v", matches)
				files = append(files, matches...)
			}
		}
		return files
	}
	return args
}

// updateFrequencies function opens each supplied file and hands it over
// to another function for the actual work/processing
func updateFrequencies(filename string, fw map[string]int) {
	logger := log.WithFields(log.Fields{
		"updateFrequencies": "word-frequencies",
	})
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open(filename); err != nil {
		logger.Errorf("unable to open the file %s %s", filename, err.Error())
		return
	}

	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), fw)
}

// readAndUpdateFrequencies function takes each file reader and a map
// of frequencies per word and then updates the map by reading the file
func readAndUpdateFrequencies(r *bufio.Reader, fw map[string]int) {
	logger := log.WithFields(log.Fields{
		"readAndUpdateFrequencies": "word-frequencies",
	})
	for {
		line, err := r.ReadString('\n')
		for _, word := range SplitAtNonLetters(strings.TrimSpace(line)) {
			// check that the word has atleast 2 characters.
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				fw[strings.ToLower(word)]++
			}
		}
		if err != nil {
			if err != io.EOF {
				logger.Errorf("failed to finish reading the file: %s", err.Error())
			}
		}
		break
	}
}

// SplitAtNonLetters function just splits a string at the non-word characters
func SplitAtNonLetters(s string) []string {
	nonLetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FieldsFunc(s, nonLetter)
}

// reportByWords function renders the data populated inside the map
// fw which has words to frequency mappings alphabetically
func reportByWords(fw map[string]int) {
	words := make([]string, 0, len(fw))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range fw {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if fwidth := len(fmt.Sprint(frequency)); fwidth > frequencyWidth {
			frequencyWidth = fwidth
		}
	}
	sort.Strings(words)
	gap := (wordWidth + frequencyWidth) - (len("Word") + len("Frequency"))
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, fw[word])
	}
}

// invertMap function inverts a map from string to int into a map
// pointing from int to []string
func invertMap(intPerString map[string]int) map[int][]string {
	stringsPerInt := make(map[int][]string, len(intPerString))
	for key, value := range intPerString {
		stringsPerInt[value] = append(stringsPerInt[value], key)
	}
	return stringsPerInt
}

// reportByFrequency function prints the output of map to Stdout by
// sorting the same by frequencies
func reportByFrequency(wordsPerFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsPerFrequency))
	for frequency := range wordsPerFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsPerFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
