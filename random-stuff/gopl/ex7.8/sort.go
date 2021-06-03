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
// printTracks function prints the track and album information
func printTracks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	tw.Flush()
}

//!-printTracks

//!+customSort
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
	swap func(x, y *Track)
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

//!+clickable
func clickable(s string) {
	switch s {
	case "title":
		sort.Sort(customSort{tracks,
			func(x, y *Track) bool {
				return x.Title < y.Title
			},
			func(x, y *Track) {
				x.Title, y.Title = y.Title, x.Title
			}})
	case "artist":
		sort.Sort(customSort{tracks,
			func(x, y *Track) bool {
				return x.Artist < y.Artist
			},
			func(x, y *Track) {
				x.Artist, y.Artist = y.Artist, x.Artist
			}})
	case "album":
		sort.Sort(customSort{tracks,
			func(x, y *Track) bool {
				return x.Album < y.Album
			},
			func(x, y *Track) {
				x.Album, y.Album = y.Album, x.Album
			}})
	case "year":
		sort.Sort(customSort{tracks,
			func(x, y *Track) bool {
				return x.Year < y.Year
			},
			func(x, y *Track) {
				x.Year, y.Year = y.Year, x.Year
			}})
	case "length":
		sort.Sort(customSort{tracks,
			func(x, y *Track) bool {
				return x.Length < y.Length
			},
			func(x, y *Track) {
				x.Length, y.Length = y.Length, x.Length
			}})

	}
}

//!-clickable

func main() {
	fmt.Println("Multi-tiered sorting")
	fmt.Printf("%s\n", strings.Repeat("*", 84))
	printTracks(tracks)

	fmt.Printf("%s\n", strings.Repeat("*", 84))
	clickable("title")
	printTracks(tracks)

	fmt.Printf("%s\n", strings.Repeat("*", 84))
	clickable("artist")
	printTracks(tracks)

	fmt.Printf("%s\n", strings.Repeat("*", 84))
	clickable("album")
	printTracks(tracks)

	fmt.Printf("%s\n", strings.Repeat("*", 84))
	clickable("year")
	printTracks(tracks)

	fmt.Printf("%s\n", strings.Repeat("*", 84))
	clickable("length")
	printTracks(tracks)
}
