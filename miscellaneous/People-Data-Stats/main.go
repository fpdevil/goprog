/*
People Database Stats (github.com/striversity/glft/)
Public Sanple Data available at https://www.mockaroo.com/

Given a file with  Comma-Separated-Values represensenting information about
individuals, write a Go program that does the following:

1. Read  the records  form the  data file. Data  file name  is passed  as a
program argument to the program.

- Hint: Use os.Args
- TIP: Examine a few lines of the data file to see what the records look like.

2. For each line read from the file, initialize a Person struct object.

- TIP: Write a function which takes a CSV string and returns a Person object.
- TIP: You will need to use the packages 'strings' and 'strconv'.
- HINT: If using the 'input.FileReader' object, be sure to check for io.EOF
  when reading records and handle it occordingly.

3. Compute the following stats by gender:

(a) total number of records
(b) min salary
(c) max salary
(d) average salary

*/

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/fpdevil/goprog/miscellaneous/People-Data-Stats/stats"

	"github.com/fpdevil/goprog/miscellaneous/People-Data-Stats/parse"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("-- People Database Stats --")

	if len(os.Args) == 1 {
		usage()
		return
	}

	// check the arguments for the csv input file
	data := os.Args[1]
	_, err := os.Stat(data)
	if os.IsNotExist(err) {
		log.Errorf("input csv data file %s is unavailable", data)
		return
	}

	// open csv file for parsing
	csv, err := os.Open(data)
	if err != nil {
		log.Error(err)
		return
	}

	// defer closing the csv file
	defer csv.Close()

	// get an instance of the Persons
	p := parse.GetPerson()

	// Read csv and assign to persons for processing
	persons, err := parse.ReadCSV(csv, p)

	// for _, v := range persons {
	// 	fmt.Printf("%v\n", v)
	// }

	pmap := parse.FillPersons(persons)

	stats.Show(pmap)
}

// usage function handles the case where the arguments are not provided
func usage() {
	_, file := path.Split(os.Args[0])
	s1 := fmt.Sprintf("Usage: go run %v <csv_file>", file)
	s2 := fmt.Sprintf("       go run . <csv_file>")
	log.Errorf(s1)
	log.Errorf(s2)
}
