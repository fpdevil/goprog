package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// number of workers which is equal to the number of processors or cores
var workers = runtime.NumCPU()

// Task represents a  specific task done over each file  and it specifies
// the file  name to be  processed and the  channel to which  any grepped
// results are sent to
type Task struct {
	filename string        // filename being processed
	result   chan<- Result // send only channel to send grepped result to
}

// Result represents  the return  of the  grep ran  over files  with the
// details of matching line and line number of a specific file.
type Result struct {
	filename string // filename of grepped result
	lino     int    // matching line number from file
	line     string // matching line from file
}

func (task Task) Do(lirx *regexp.Regexp) {
	file, err := os.Open(task.filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lirx.Match(line) {
			task.result <- Result{task.filename, lino, string(line)}
		}
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "error:%d: %s\n", lino, err)
			}
			break
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // using all the available cores
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <files>\n", filepath.Base(os.Args[0]))
		return
	}

	lirx, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid regular expression: %s\n", err)
		return
	}

	grep(lirx, parseCmdLineArgs(os.Args[2:]))
}

func parseCmdLineArgs(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // this is an invalid pattern
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func grep(lirx *regexp.Regexp, files []string) {
	tasks := make(chan Task, workers)
	results := make(chan Result, minimum(1000, len(files)))
	done := make(chan struct{}, workers)

	go addTasks(tasks, files, results) // this executes in its own goroutine
	for i := 0; i < workers; i++ {
		go doTasks(done, lirx, tasks) // executes in its own goroutine
	}
	go awaitCompletion(done, results)
	processResults(results)
}

func addTasks(tasks chan<- Task, files []string, results chan<- Result) {
	for _, file := range files {
		tasks <- Task{filename: file, result: results}
	}
	close(tasks)
}

func doTasks(done chan<- struct{}, lirx *regexp.Regexp, tasks <-chan Task) {
	for task := range tasks {
		task.Do(lirx)
	}
	done <- struct{}{}
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("* %s:%d:%s\n", result.filename, result.lino, result.line)
	}
}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}
