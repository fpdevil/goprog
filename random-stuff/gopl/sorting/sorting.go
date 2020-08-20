package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

// Track represents data from musical albums
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"2045: Radical Man", "Prince", "Bamboozled", 2000, length("6m36s")},
	{"The Calm", "Andy Allo", "Superconductor", 2012, length("5m20s")},
	{"Cold Coffee & Cocaine", "Prince", "Piano & A Microphone", 1983, length("4m17s")},
}

//!+printTracks
// printTracks function prints the playlist as a table
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	// initialize the tabwriter
	tw := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	tw.Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

//!-printTracks

//!+artist
// byArtist type for sorting the tracks by Artist name
type byArtist []*Track

// Len is the number of elements in the collection.
func (a byArtist) Len() int { return len(a) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (a byArtist) Less(i, j int) bool {
	return a[i].Artist < a[j].Artist
}

// Swap swaps the elements with indexes i and j.
func (a byArtist) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//!-artist

//!+year
// byYear type for sorting by Year
type byYear []*Track

// Len is the number of elements in the collection.
func (a byYear) Len() int { return len(a) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (a byYear) Less(i, j int) bool {
	return a[i].Year < a[j].Year
}

// Swap swaps the elements with indexes i and j.
func (a byYear) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//!-year

//!+customSort
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// Len is the number of elements in the collection.
func (a customSort) Len() int { return len(a.t) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (a customSort) Less(i, j int) bool {
	return a.less(a.t[i], a.t[j])
}

// Swap swaps the elements with indexes i and j.
func (a customSort) Swap(i, j int) {
	a.t[i], a.t[j] = a.t[j], a.t[i]
}

//-customSort

func main() {
	fmt.Println("Sort By Artist")
	fmt.Printf("%s\n", strings.Repeat("*", 75))
	// first convert tracks to new type and apply sorting
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Sort By Artist in Reverse")
	fmt.Printf("%s\n", strings.Repeat("*", 75))
	// first convert tracks to new type and apply sorting
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Sort By Year")
	fmt.Printf("%s\n", strings.Repeat("*", 75))
	// first convert tracks to new type and apply sorting
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Multi-tier ordering")
	fmt.Printf("%s\n", strings.Repeat("*", 75))
	// define a multi-tier ordering function whose primary sort key
	// is Title, whose secondary key is Year, and whose tertiary key
	// is the running time, length
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}
