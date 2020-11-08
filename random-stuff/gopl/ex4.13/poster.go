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
http://www.omdbapi.com/?apikey=1a678916&t=interstellar
*/

// Movie struct contains selected items of interest from the omdbapi movie
// search results json
type Movie struct {
	Title    string
	Year     string
	Poster   string
	Response string
	Error    string
}

const (
	// MovieAPI points to the base url of omdbapi
	MovieAPI = `https://omdbapi.com/?apikey=%s&t=%v`
	// usage details string
	usage = `
	go run poster.go <apikey> <moviename>
	`
)

// findMovie function finds the specified movie title in the omdbapi website
// and returns the movie details, else an error message with appropriate details
// the service is protecetd with an apikey, which needs to be passed as an argument
func findMovie(apikey string, title []string) (movie Movie, err error) {
	q := url.QueryEscape(strings.Join(title, " "))
	url := fmt.Sprintf(MovieAPI, apikey, q)
	log.Infof("calling movie db @ %v", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calling %s: %v", url, err)
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
		log.Errorf("error decoding movie data: %v", err)
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
		log.Errorf("Usage: %s", usage)
		return
	}

	apikey := args[0]
	title := args[1:]

	movie, err := findMovie(apikey, title)
	if err != nil {
		log.Errorf("error looking for movie title: %v", title)
		return
	}

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
