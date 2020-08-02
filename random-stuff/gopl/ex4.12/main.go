package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	usage = `
		xkcd get id
	`
	xkcdurl = `https://xkcd.com/%d/info.0.json`
)

// Comic struct represents the json mapping for go struct
type Comic struct {
	Num        int
	Year       string
	Month      string
	Day        string
	Link       string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Title      string
}

// fetchComic function fetches the comic matching specified
// id in the argument from xkcd url
func fetchComic(id int) (*Comic, error) {
	var comic Comic
	url := fmt.Sprintf(xkcdurl, id)
	log.Infof("calling url %v", url)

	res, err := http.Get(url)
	if err != nil {
		log.Errorf("error (%v) calling %s", err, url)
		return &comic, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return &comic, fmt.Errorf("unable to fetch comic %d status: %s", id, res.StatusCode)
	}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&comic); err != nil {
		log.Errorf("%v decoding json data", err)
		return &comic, err
	}

	return &comic, nil
}

// ParseDate function returns a string representation of the comic date
func (c *Comic) ParseDate() string {
	return fmt.Sprintf("%s-%s-%s", c.Day, c.Month, c.Day)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stdout, "please provide a valid input data")
		return
	}

	cmd := args[1]
	n, _ := strconv.Atoi(cmd)
	comic, err := fetchComic(n)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	fmt.Printf("%v\n", comic)
}
