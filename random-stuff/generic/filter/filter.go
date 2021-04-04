package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // will use all available cores of machine
	algorithm, minsize, maxsize, suffixes, dirs := handleCmdLineArgs()

	if algorithm == 1 {
		sink(filterSize(minsize, maxsize, filterSuffixes(suffixes, source(listFiles(dirs)))))
	} else {
		channel1 := listFiles(dirs)
		channel2 := source(channel1)
		channel3 := filterSuffixes(suffixes, channel2)
		channel4 := filterSize(minsize, maxsize, channel3)
		sink(channel4)
	}
}

func handleCmdLineArgs() (algorithm int, minsize, maxsize int64, suffixes, directories []string) {
	flag.IntVar(&algorithm, "algorithm", 1, "should be 1 or 2")
	flag.Int64Var(&minsize, "min", -1, "minimum file size [-1 means no minimum]")
	flag.Int64Var(&maxsize, "max", -1, "maximum file size [-1 means no maximum]")

	var dirs *string = flag.String("directories", ".", "comma delimited list of directory locations")
	var suffixList *string = flag.String("suffixes", "", "comma delimited list of file suffixes")
	flag.Parse()

	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}

	if minsize > maxsize && maxsize != -1 {
		fmt.Printf("minimum size should be < maximum size")
		return
	}

	suffixes = []string{}
	if *suffixList != "" {
		suffixes = strings.Split(*suffixList, ",")
	}

	directories = []string{}
	if *dirs != "" {
		directories = strings.Split(*dirs, ",")
	}

	// files = flag.Args()
	// return algorithm, minsize, maxsize, suffixes, files
	return algorithm, minsize, maxsize, suffixes, directories
}

func listFiles(dirs []string) <-chan string {
	out := make(chan string, 10)
	go func() {
		for _, loc := range dirs {
			files, err := walkDir(loc)
			if err != nil {
				continue
			}
			for _, f := range files {
				out <- filepath.Join(loc, f)
			}
		}
		close(out)
	}()
	return out
}

func source(in <-chan string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for filename := range in {
			out <- filename
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			extn := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if extn == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func filterSize(minimum, maximum int64, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if minimum == -1 && maximum == -1 {
				out <- filename
				continue
			}

			finfo, err := os.Stat(filename)
			if err != nil {
				fmt.Printf("error %v\n", err)
				continue
			}

			size := finfo.Size()
			if (minimum == -1 || minimum > -1 && minimum <= size) &&
				(maximum == -1 || maximum > -1 && maximum >= size) {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func walkDir(root string) ([]string, error) {
	var result []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, err
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Printf("%v\n", filename)
	}
}
