package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Concurrent File filtering process...")
	fmt.Println()

	// we will use all the available cores of cpu
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetFormatter(&log.TextFormatter{})
	algorithm, minSize, maxSize, suffixes, files := parseCmdLineArgs()

	if algorithm == 1 {
		sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files))))
	} else {
		channel1 := source(files)
		channel2 := filterSuffixes(suffixes, channel1)
		channel3 := filterSize(minSize, maxSize, channel2)

		sink(channel3)
	}
}

// parseCmdLineArgs function handles the command line arguments supplied
func parseCmdLineArgs() (algorithm int, minSize, maxSize int64, suffixes, files []string) {
	flag.IntVar(&algorithm, "algorithm", 1, "specify algorithm to be 1 or 2")
	flag.Int64Var(&minSize, "min", -1, "minimum file size [if -1 then no mimimum]")
	flag.Int64Var(&maxSize, "max", -1, "maximum file size [if -1 then no maximum]")
	var suffixOps *string = flag.String("suffixes", "", "comma separated list of file suffixes for filtering")
	flag.Parse()

	// if algorithm is beyond 1 or 2, default to a value 1
	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}

	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minimum size must be less than maximum size")
	}

	suffixes = []string{}
	if *suffixOps != "" {
		suffixes = strings.Split(*suffixOps, ",")
	}

	files = flag.Args()
	return algorithm, minSize, maxSize, suffixes, files
}

// source function creates a channel for passing the filenames listed
func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		defer close(out)
		for _, filename := range files {
			out <- filename
		}
	}()
	return out
}

// filterSuffixes function creates a buffer of same size as files
// and performs filtering based in the input suffixes if there are
// no suffixes provided, then all the files from input channel will
// be sent as is to the output for next channel processing
func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		defer close(out)
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}

			ext := strings.ToLower(filepath.Ext(filename))
			// if the filename extension matches the one in suffixes, then
			// we will push it to the out channel else go for next iteration
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
	}()
	return out
}

// filterSize function filters the files based on their size specifications
func filterSize(minimum, maximum int64, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		defer close(out)
		for filename := range in {
			if minimum == -1 && maximum == -1 {
				out <- filename
				continue
			}

			fileinfo, err := os.Stat(filename)
			if err != nil {
				continue
			}

			size := fileinfo.Size()
			if (minimum == -1 || minimum > -1 && minimum <= size) &&
				(maximum == -1 || maximum > -1 && maximum >= size) {
				out <- filename
			}
		}
	}()
	return out
}

// sink function iterates over all the filenames from the last channel
// in the main goroutine and will mark the closure
func sink(in <-chan string) {
	for filename := range in {
		fmt.Printf("* %v\n", filename)
	}
}
