package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	// xkcdurl is the url to fetch the comic based on id
	xkcdurl = `https://xkcd.com/%s/info.0.json`
)

// Index struct creates an index of the xkcd comics
type Index struct {
	Comics   map[int]*Comic
	FilePath string
}

// Comic struct represents the json mapping for go struct
type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// xkcdURL function returns a url for the specific comic
// metadata based on the comic id provided
func xkcdURL(comicID int) string {
	return fmt.Sprintf(xkcdurl, strconv.Itoa(comicID))
}

// matches function checks if two strings have common terms
// by doing a case insensitive lookup
func matches(str1, str2 string) bool {
	return strings.Contains(strings.ToLower(str1), strings.ToLower(str2))
}

// FetchComic function fetches a single comic matching specified
// id in the argument from xkcdurl url
func FetchComic(id int) (*Comic, error) {
	var comic Comic
	url := xkcdURL(id)
	log.Infof("calling url %v", url)

	res, err := http.Get(url)
	if err != nil {
		log.Errorf("error (%v) calling %s", err, url)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Errorf("http status %d unable to fetch the coming %d", res.StatusCode, id)
		return nil, fmt.Errorf("unable to fetch comic %v", id)
	}

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&comic); err != nil {
		log.Errorf("%v decoding json data", err)
		return nil, err
	}

	return &comic, nil
}

// New function checks for an existing index file at the location
// specified in filepath and if it does not exist populates one
func New(filepath string) *Index {
	var comics map[int]*Comic

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "New",
			"file":     filepath,
		}).Errorf("error while reading index %s\n", err.Error())
		comics = make(map[int]*Comic)
	} else {
		json.Unmarshal(data, &comics)
	}

	return &Index{
		Comics:   comics,
		FilePath: filepath,
	}
}

// insert function inserts a comic from the map into the index
func (index *Index) insert(comics map[int]*Comic) {
	log.WithFields(log.Fields{
		"function": "insert",
	}).Debugf("adding comic id from map to index")
	for id := range comics {
		index.Comics[id] = comics[id]
	}
}

// Build function builds an index between the intervals specified
// inclusive of from and to comic id's and populates it in Index
func (index *Index) Build(from, to int) {
	comics := make(map[int]*Comic)
	for i := from; i < to; i++ {
		// if map already contains comic skip the loop for next
		c, ok := index.Comics[i]
		if ok {
			log.WithFields(log.Fields{
				"function": "Build",
				"comic":    c,
			}).Debug("comic exists in the map, continuing for next")
			continue
		}

		// fetch comic from remote and update the map
		comic, err := FetchComic(i)
		if err != nil {
			log.WithFields(log.Fields{
				"function": "Build",
				"comic id": i,
			}).Error("error fetcing the comic")
			continue
		}

		comics[i] = comic
	}
	index.insert(comics)
}

// Save function persists the index data to the file
func (index *Index) Save() {
	data, err := json.MarshalIndent(index.Comics, "", "\t")
	if err != nil {
		log.WithFields(log.Fields{
			"function": "Save",
			"error":    err.Error(),
		}).Error("error while marshalling the comic data")
		return
	}

	ioutil.WriteFile(index.FilePath, data, 0644)
}

// Search function searches for the query term in the persisted indexed file
func (index *Index) Search(term string) []*Comic {
	var comics []*Comic
	for _, i := range index.Comics {
		if matches(i.Title, term) ||
			matches(i.Transcript, term) ||
			matches(i.SafeTitle, term) {
			comics = append(comics, i)
		}
	}
	return comics
}

// Add function adds an index entry about comic into index file
func (index *Index) Add() {
	data, err := json.MarshalIndent(index.Comics, "", "\t")
	if err != nil {
		log.WithFields(log.Fields{
			"function": "Add",
			"error":    err.Error(),
		}).Error("error marshalling json data")
		return
	}
	ioutil.WriteFile(index.FilePath, data, 0644)
}

// ParseDate function returns a string representation of the comic date
func (c *Comic) ParseDate() string {
	return fmt.Sprintf("%s-%s-%s", c.Day, c.Month, c.Day)
}
