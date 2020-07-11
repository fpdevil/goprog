package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Cities struct will hold the main data fields from csv file
type Cities struct {
	Rank       int
	City       string
	Country    string
	Population uint
}

// CityMap is a hashmap from the each country to the Cities struct
type CityMap map[string][]Cities

// newCity function returns an empty struct Cities{} for initializing
// the processing
func newCity() Cities {
	return Cities{}
}

// newCities function returns an empty list of struct Cities{} for
// initializing the processing
func newCities() []Cities {
	c := []Cities{}
	return c
}

// newCityMap returns a map from each country to all fields of each
// matching record to that country
func newCityMap() CityMap {
	return make(map[string][]Cities)
}

// fillCities function populates the map with necessary filtered ddata
func fillCities(cities []Cities) (cMap CityMap) {
	cMap = newCityMap()
	for _, v := range cities {
		cMap[v.Country] = append(cMap[v.Country], v)
	}
	return
}

// getCountries function picks the countries with largest number of
// states as specified and sorts them and returns
func getCountries(states int, cmap CityMap) []string {
	countries := make([]string, 0)

	for _, v := range cmap {
		if len(v) >= states {
			countries = append(countries, v[0].Country)
			// fmt.Printf("%s: %v\n", k, v)
		}
	}
	// sort the list of countries
	sort.Strings(countries)

	return countries
}

// comparator is the helper function aiding in the filtering process
// of the custom Cities type structure
func comparator(c int) func([]Cities) bool {
	return func(cities []Cities) bool {
		return len(cities) >= c
	}
}

// parseAllCSV function reads and parses the entire CSV file
// at once and might not be much efficient
func parseAllCSV(c []Cities, d io.Reader) {
	// initialize the reader
	rr := csv.NewReader(bufio.NewReader(d))

	// read all the records at once
	records, err := rr.ReadAll()

	if err != nil {
		log.Errorf("error occurred while parsing: %q\n", err)
		return
	}

	// Iterate through the read records
	for _, v := range records {
		rec := Cities{
			City:    v[1],
			Country: v[2],
		}
		c = append(c, rec)
	}

	for _, v := range c {
		log.Printf("%v\n", v)
	}
}

// parseCsvByRow parses the given csv file row by row which
// is more efficient
func parseCsvByRow(c Cities, d io.Reader) ([]Cities, error) {

	var (
		cities []Cities // list of cities
		check  bool     // a predicate for filtering csv header
		cline  int      // current line number from csv
	)

	rr := csv.NewReader(bufio.NewReader(d))

	header, err := rr.Read()
	if err != nil {
		log.Error("an error encountered while reading ::", err)
		return nil, err
	}
	log.Printf("csv headers from the data file: %v \n", header)

	for {
		cline++
		record, err := rr.Read()
		if err == io.EOF {
			// log.Error("EOF error encountered while reading ::", err)
			break
		} else if err != nil {
			log.Error("an error encountered while reading ::", err)
			return nil, err
		}

		// fmt.Printf("Row %d : %v \n", cline, record)
		// starting with the check=false will ensure that we
		// are essentially skipping the csv headers
		if check {
			for i, val := range record {
				switch i {
				case 0:
					val = strings.TrimSpace(val)
					c.Rank, err = strconv.Atoi(val)
					if err != nil {
						log.Errorf("error processing record :: %d\n", cline)
						log.Error(err)
						return nil, err
					}
				case 1:
					c.City = val
				case 2:
					c.Country = val
				case 3:
					pop, err := strconv.ParseUint(val, 10, 64)
					c.Population = uint(pop)
					if err != nil {
						log.Errorf("error processing record :: %d\n", cline)
						log.Error(err)
						return nil, err
					}
				}
			}
			cities = append(cities, c)
		}
		check = true
	}
	// fmt.Printf("%v\n", cities)
	return cities, nil
}
