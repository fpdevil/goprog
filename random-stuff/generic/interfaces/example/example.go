package main

import (
	"fmt"
	"time"
)

type Entry interface {
	Title() string
}

type Book struct {
	Name      string
	Author    string
	Published time.Time
}

type Movie struct {
	Name     string
	Director string
	Year     int
}

func (b Book) Title() string {
	return fmt.Sprintf("%s by %s (%s)", b.Name, b.Author, b.Published.Format("Jan 2006"))
}

func (m Movie) Title() string {
	return fmt.Sprintf("%s (%d)", m.Name, m.Year)
}

func Display(e Entry) string {
	return e.Title()
}

func main() {
	b := Book{
		Name:      "20 Thousand leagues under the sea",
		Author:    "Jules Verne",
		Published: time.Date(1861, time.October, 21, 0, 0, 0, 0, time.UTC),
	}

	m := Movie{
		Name:     "The Guns Of Navarone",
		Director: "J. Lee Thompson",
		Year:     1961,
	}

	fmt.Println(Display(b))
	fmt.Println(Display(m))
}
