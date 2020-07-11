package parse

import (
	"bufio"
	"encoding/csv"
	"io"
	"strconv"

	"github.com/fpdevil/goprog/miscellaneous/People-Data-Stats/types"

	log "github.com/sirupsen/logrus"
)

// GetPerson gets an instance of the Person struct and returns and empty
// structure for data
func GetPerson() types.Person {
	return types.Person{}
}

// GetPersons getsan instance of the Persons struct which is a group of
// persons together providing details of all persons
func GetPersons() types.Persons {
	return types.Persons{}
}

// GetPMap gives an instance of the map
func GetPMap() types.PMap {
	return make(map[rune]types.Persons)
}

// ReadCSV is the main parsing function which handles the parsing of
// the input csv file and populates data from the csv into Person
func ReadCSV(d io.Reader, p types.Person) (types.Persons, error) {

	var (
		persons types.Persons
		cline   int
		track   bool
	)
	reader := csv.NewReader(bufio.NewReader(d))

	header, err := reader.Read()
	if err != nil {
		log.Errorf("error occurred parsing csv headers :: %s", err)
		return nil, err
	}
	log.Printf("csv headers from the data file: %v \n", header)

	for {
		cline++
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Errorf("error occurred parsing csv :: %s", err)
			return nil, err
		}

		if track {
			for i, val := range record {
				switch i {
				case 0:
					p.FirstName = val
				case 1:
					p.LastName = val
				case 2:
					p.SSN = val
				case 3:
					p.Gender = []rune(val)[0] // converting string to rune
				case 4:
					age, err := strconv.ParseUint(val, 10, 8)
					if err != nil {
						log.Errorf("error parsing csv at line %d :: %s", cline, err)
						return nil, err
					}
					p.Age = uint8(age)
				case 5:
					salary, err := strconv.ParseFloat(val, 10)
					if err != nil {
						log.Errorf("error parsing csv at line %d :: %s", cline, err)
						return nil, err
					}
					p.Salary = types.Currency(salary)
				}
			}
			persons = append(persons, p)
		}
		track = true
	}
	return persons, err
}

// FillPersons function populates the map with data
func FillPersons(persons types.Persons) (pmap types.PMap) {
	pmap = GetPMap()
	for _, p := range persons {
		pmap[p.Gender] = append(pmap[p.Gender], p)
	}
	return pmap
}
