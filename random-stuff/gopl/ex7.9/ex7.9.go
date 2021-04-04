package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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

// tracks variable contains a playlist
// each track is a single row and each column is an attribute of the
// track like title, artist, album, year and duration; each element
// is indirectly a pointer to a track
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"2045: Radical Man", "Prince", "Bamboozled", 2000, length("6m36s")},
	{"The Calm", "Andy Allo", "Superconductor", 2012, length("5m20s")},
	{"Cold Coffee & Cocaine", "Prince", "Piano & A Microphone", 1983, length("4m17s")},
	{"Feelin' the Same Way", "Norah Jones", "Come Away with Me", 2002, length("2m57s")},
	{"Find the Answer", "Arashi", "5x20 All the Best!! 1999–2019", 2020, length("4m3s")},
}

//!+length

// length function returns track duration in appropriate seconds
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-length

var tracktab = template.Must(template.New("tracktab").Parse(`
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>Tracks Table</title>
    </head>
    <body>
        <h1>Tracks</h1>
        <table>
            <thead>
                <tr>
                    <th><a href="/?sort=title">Title</a></th>
                    <th><a href="/?sort=Artist">Artist</a></th>
                    <th><a href="/?sort=Album">Album</a></th>
                    <th><a href="/?sort=Year">Year</a></th>
                    <th><a href="/?sort=Length">Length</a></th>
                </tr>
            </thead>
            <tbody>
                {{range .}}
                <tr>
                    <td>{{.Title}}</td>
                    <td>{{.Artist}}</td>
                    <td>{{.Album}}</td>
                    <td>{{.Year}}</td>
                    <td>{{.Length}}</td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </body>
</html>
`))

//!+exec

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "error serving page: %s", err.Error())
		return
	}
}

//!-exec

//!+handler

//!-handler

//!+printTracks
// printTracks function prints the playlist as a table
func printTable(w io.Writer, tracks []*Track) {
	if err := tracktab.Execute(w, tracks); err != nil {
		fmt.Fprintf(os.Stderr, "error writing %s", err.Error())
		return
	}
}

//!-printTracks

//!+artist
// byArtist type for sorting the tracks by Artist name
type byArtist []*Track

// Len is the number of elements in the collection.
func (a byArtist) Len() int {
	return len(a)
}

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
