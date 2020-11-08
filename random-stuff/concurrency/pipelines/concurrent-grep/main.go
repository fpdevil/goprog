package main

/*
 * The concurrent  grep program take a  regular expression and a  list of
 * files on  the command line and  output the filename, line  number, and
 * every  line in  every file  where the  regular expression  matches. If
 * there are no matches, then there is no output.
 */

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	Formatter := new(log.JSONFormatter)
	log.SetFormatter(Formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// Result represents the structure of the response from running grep
type Result struct {
	filename string // the name of the file
	line     string // the matching line from the filename
	lineno   int    // the corresponding matching line number
}

// Job represents the current job/task of grepping and sending results
type Job struct {
	filename string        // name of the file to be processed
	results  chan<- Result // sender channel to which results would be sent
}

const (
	msg = `* Concurrent Grep over Files *

	`
	usage = `%s <regexp> <files>\n`
)

var (
	// number of worker threads same as cpu's
	workers = runtime.NumCPU()
	// initialize the logger
	logger = log.WithFields(log.Fields{
		"main": "concurrent grep",
	})
)

func main() {
	fmt.Println(msg)

	// use all the available cpu cores of machine
	runtime.GOMAXPROCS(runtime.NumCPU())
	args := os.Args
	if len(args) < 3 || args[1] == "-h" || args[1] == "--help" {
		fmt.Fprintf(os.Stderr, usage, filepath.Base(args[0]))
		return
	}
	if lineRx, err := regexp.Compile(args[1]); err != nil {
		logger.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, parseCmdLineArgs(args[2:]))
	}
}

func parseCmdLineArgs(files []string) []string {
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

// grep function 3  bidirectional channels as needed by  the program. The
// task or  jobs are  the spread  over as many  worker goroutines  as the
// number of processors available. The buffer  size will be equal to this
// number to avoid blocking.
func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	// add jobs to the job channel
	go addJobs(jobs, filenames, results)
	for i := 0; i < workers; i++ {
		// perform work
		go doJobs(done, lineRx, jobs)
	}
	go awaitCompletion(done, results)

	processResults(results)
}

// addJobs function sends every filename to the jobs channel, one by one
// as a job value.  In short the function starts adding jobs to the jobs
// channel via a goroutine.
func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

// doJobs function is  invoked for the NumCPU() times to  share the work.
// Each invocation iterates over the  same shared jobs channel (a receive
// only channel), and each will be  blocked until a job becomes available
// for it to receive.
func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	// indicate running out of job by sending an empty struct
	done <- struct{}{}
}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}

func (job Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for lineno := 1; ; lineno++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {
			job.results <- Result{job.filename, string(line), lineno}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("error: %d: %s", lineno, err)
			}
			break
		}
	}
}

// awaitCompletion function  ensures that the main  goroutine waits until
// all the processing is done before terminating, thus avoiding the below
// two pitfals.
// 1. When the program finishes  almost immediately, but does not produce
// results.  This happens  because when  main go  routine terminates  the
// other goroutines processing will all be terminated
// 2. Avoid deadlocks
func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

// processResults function iterates over the results channel for results
func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s: %d:%s\n", result.filename, result.lineno, result.line)
	}
}

// waitAndProcessResults
func waitAndProcessResults(timeout int64, done <-chan struct{}, results <-chan Result) {
	finish := time.After(time.Duration(timeout))
	for working := workers; working > 0; {
		select {
		case result := <-results:
			fmt.Printf("%s: %d:%s\n", result.filename, result.lineno, result.line)
		case <-finish:
			fmt.Println("timed out")
			return
		case <-done:
			working--
		}
	}
	for {
		select {
		case result := <-results:
			fmt.Printf("%s: %d:%s\n", result.filename, result.lineno, result.line)
		case <-finish:
			fmt.Println("timed out")
			return
		default:
			return
		}
	}
}
