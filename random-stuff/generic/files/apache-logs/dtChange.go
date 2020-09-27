// Parse and change date time of apache web server logs
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: %v <input file>\n", filepath.Base(args[0]))
		return
	}

	filename := args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file %s\n", err)
		return
	}
	defer file.Close()

	// number of lines in the input file which do not match any
	// of the regular expressions
	var nomatch int

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file: %s", err)
		}

		// target string: [21/Nov/2017:19:28:09 +0200]
		// format sring:  Mon Jan 2 15:04:05 -0700 MST 2006
		str1 := `.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`
		re1 := regexp.MustCompile(str1)
		// fmt.Println(re1.MatchString(line))
		if re1.MatchString(line) {
			match := re1.FindStringSubmatch(line)
			dt1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := dt1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				nomatch++
			}
			continue
		}

		// target string: [Jun-21-17:19:28:09 +0200]
		// format sring:  Mon Jan 2 15:04:05 -0700 MST 2006
		str2 := `.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`
		re2 := regexp.MustCompile(str2)
		if re2.MatchString(line) {
			match := re2.FindStringSubmatch(line)
			dt2, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := dt2.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				nomatch++
			}
			continue
		}
	}
	fmt.Printf("%v lines did not match!\n", nomatch)
}
