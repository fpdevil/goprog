package types

import (
	"fmt"
	"strings"
)

// Currency is the type for representing salary
type Currency float64

// Person struct represents the person entity from csv file
type Person struct {
	FirstName string
	LastName  string
	SSN       string
	Gender    rune
	Age       uint8
	Salary    Currency
}

// Persons represent a group of persons
type Persons []Person

// PMap holds the statistics by gender
type PMap map[rune]Persons

// Stats will hold the details of person stats
type Stats struct {
	Total     int
	MinSalary Currency
	MaxSalary Currency
	AvgSalary Currency
}

// String() for Currency formatting display
func (c Currency) String() string {
	return fmt.Sprintf("$%1.2f", float64(c))
}

func (p Persons) Len() int           { return len(p) }
func (p Persons) Less(i, j int) bool { return p[i].Salary < p[j].Salary }
func (p Persons) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Display the statistics
func (s Stats) Display(d string) {
	fmt.Printf("%s\n", strings.Repeat("=", 36))
	fmt.Printf("%v Statistics:\n", d)
	fmt.Printf("%s\n", strings.Repeat("=", 36))
	fmt.Printf("Count		: %-10v\n", s.Total)
	fmt.Printf("Min Salary	: %-10v\n", s.MinSalary)
	fmt.Printf("Max Salary	: %-10v\n", s.MaxSalary)
	fmt.Printf("Avg Salary	: %-10v\n", s.AvgSalary)
	fmt.Printf("%s\n", strings.Repeat("-", 36))
	fmt.Println()
}
