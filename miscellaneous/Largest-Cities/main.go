package main

/*
Data from http://www.citymayors.com/features/largest_cities_3.html

Given a CSV file representing the largest cities in the World, write a
`GO` program which does the following:

Read the records form the data file.  Data file name is passed as an
input argument to the program.

- Hint: Use os.Args
- TIP:  Examine a few lines of the data file to see what the records
	    look like.

For each line read from the file, initialize a struct object.
- TIP	: Write a function which takes a `CSV` string and returns a object.
- TIP	: You will need to use the packages 'strings' and 'strconv'.

Print the names of the countries with 5 or more of the largest cities
in the data set.

Requirements:
- You must use a Struct type to represent a record from the file.
- Use more than one files for readability.


	running: `go run . cities.csv`
*/

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// countries slice to hold countries with 5 or more largest cities
var n = 5

func main() {
	log.Println("/// Largest Cities in the World ///")

	// check the input arguments
	if len(os.Args) == 1 {
		usage()
		return
	}

	// check the arguments for the csv input file
	data := os.Args[1]
	_, err := os.Stat(data)
	if os.IsNotExist(err) {
		log.Errorf("input `csv` data file %s does not exist", data)
		return
	}

	// now open the csv file for parsing
	csv, err := os.Open(data)
	if err != nil {
		log.Error(err)
		return
	}

	// defer closing the csv file
	defer csv.Close()

	// c := newCities()
	// parseAllCSV(c, csv)

	// get an instance of the City{} struct
	city := newCity()

	// parse the csv file row by row
	s, err := parseCsvByRow(city, csv)
	if err != nil {
		log.Errorf("error reading the data :: %q\n", err)
	}

	// cities map filled in
	cMap := fillCities(s)
	countries := getCountries(n, cMap)

	fmt.Printf("Countries with %d or more of the largest cities in the world:\n", n)
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	for _, c := range countries {
		fmt.Printf("\t%v\n", c)
	}

	// g := comparator(5)
	// for _, cities := range cMap {
	// 	if g(cities) {
	// 		fmt.Printf("\t%v\n", cities[0].Country)
	// 	}
	// }
}

// usage function handles the case where the arguments are not provided
func usage() {
	s := fmt.Sprintf("Usage: %s [csv file]\n", os.Args[0])
	log.Errorf(s)
	// fmt.Fprintf(os.Stderr, "Usage: %s [csv file]\n", os.Args[0])
	// flag.PrintDefaults()
}
