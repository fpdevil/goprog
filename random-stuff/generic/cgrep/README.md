# Concurrent Grep

An implementation of the standard `unix/linux` grep command using `go` concurrency.

## Run
```sh
⇒  go run main.go "grep" README.md main.go

   main.go:  40: 		grep(lineRx, parseCmdLineFiles(os.Args[2:]))
 README.md:   3: An implementation of the standard `unix/linux` grep command using `go` concurrency.
   main.go:  88: //!+ grep
   main.go:  90: // grep function searches for the regular expression over list of files
   main.go:  91: func grep(lineRx *regexp.Regexp, filenames []string) {
```
