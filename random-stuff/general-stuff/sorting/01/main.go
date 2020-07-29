package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track is the custom type for representing playlists
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// printTracks function prints the playlist as a table using tabwriter
// package for producing the table with aligned columns
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// satisfy the interface sort.Sort for implementing the sort
type byArtist []*Track

// Len is the number of elements in the collection.
func (ba byArtist) Len() int {
	return len(ba)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (ba byArtist) Less(i, j int) bool {
	return ba[i].Artist < ba[j].Artist
}

// Swap swaps the elements with indexes i and j.
func (ba byArtist) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

// sort by year
type byYear []*Track

// Len is the number of elements in the collection.
func (by byYear) Len() int {
	return len(by)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (by byYear) Less(i, j int) bool {
	return by[i].Year < by[j].Year
}

// Swap swaps the elements with indexes i and j.
func (by byYear) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

// customSort
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// Len is the number of elements in the collection.
func (cs customSort) Len() int { return len(cs.t) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (cs customSort) Less(i, j int) bool { return cs.less(cs.t[i], cs.t[j]) }

// Swap swaps the elements with indexes i and j.
func (cs customSort) Swap(i, j int) { cs.t[i], cs.t[j] = cs.t[j], cs.t[i] }

func main() {
	fmt.Println("* sorting *")
	fmt.Println()

	fmt.Println("Sorting By Artist")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Sorting By Artist (In Reverse)")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Sorting By Year")
	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Custom Sorting")
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
