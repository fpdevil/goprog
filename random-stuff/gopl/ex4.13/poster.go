package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

/*
An example invocation of the movie search API using the api key
from omdb is as below:
http://www.omdbapi.com/?apikey=1a678916&t=guns+of+navarone
*/

// Movie struct contains selected items of interest from the omdbapi movie
// search results json
type Movie struct {
	Title    string
	Year     string
	Poster   string
	Response string
	Actors   string
	Plot     string
	Error    string
}

const (
	// MovieAPI points to the base url of omdbapi with place
	// holders for query parametes apikey and movie title
	MovieAPI = `https://omdbapi.com/?apikey=%s&t=%v`
	// a short program usage instruction string
	usage = `
	usage: go run %s <apiKey> <movie title string>

	`
)

// movieURL function returns a qualified url with proper query parameters
// for getting the movie from omdb remote url
func movieURL(apikey, query string) string {
	return fmt.Sprintf(MovieAPI, apikey, query)
}

// findMovie  function searches  for  the specified  movie  title in  the
// omdbapi website  using the  registered api key  and returns  the movie
// details. On failure, an error message with appropriate details will be
// returned. The api/service is proteceted with an apikey, which needs to
// be passed as an argument
func findMovie(apikey string, title []string) (movie Movie, err error) {
	q := url.QueryEscape(strings.Join(title, " "))
	url := movieURL(apikey, q)
	log.Infof("initiating movie title query from omdb db @ %v", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching data from %s: %v", url, err.Error())
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Errorf("movie db url @ %v received %v", url, res.StatusCode)
		err = fmt.Errorf("%s received %d response code", url, res.StatusCode)
		return
	}

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&movie); err != nil {
		log.Errorf("error decoding movie data: %v", err.Error())
		return
	}
	return
}

func (m Movie) takePosterName() string {
	extn := filepath.Ext(m.Poster)
	re := regexp.MustCompile(`\s+`)
	title := re.ReplaceAllString(strings.TrimSpace(m.Title), "_")
	return fmt.Sprintf("%s%s", title, extn)
}

func (m Movie) writePoster() (err error) {
	posterLink := m.Poster
	log.Infof("fetching the poster from %s", posterLink)
	res, err := http.Get(posterLink)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calling %s: %v", posterLink, err)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Errorf("movie db url @ %v received %v", posterLink, res.StatusCode)
		err = fmt.Errorf("%s received %d response code", posterLink, res.StatusCode)
		return
	}

	// create file for the poster
	file, err := os.Create(m.takePosterName())
	if err != nil {
		log.Errorf("error writing poster %v", err)
		return
	}
	defer file.Close()

	// now write to the created poster file
	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(res.Body)
	if err != nil {
		log.Errorf("error reading from poster %v", err)
		return
	}
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, usage, filepath.Base(os.Args[0]))
		return
	}

	apikey := args[0]
	title := args[1:]

	movie, err := findMovie(apikey, title)
	if err != nil {
		log.Errorf("error looking for movie title: %v", title)
		return
	}

	log.Infof("movie information retrieved: %#v\n", movie)
	if (Movie{} == movie) {
		log.Errorf("No results for '%s'", title)
		return
	}

	err = movie.writePoster()
	if err != nil {
		log.Errorf("error writing '%v'", err)
		return
	}
}
