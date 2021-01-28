package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/fpdevil/goprog/random-stuff/gopl/ex4.12/xkcd"
	log "github.com/sirupsen/logrus"
)

const (
	usage = `
		xkcd get id
	`
	indexDB = "index.json"
)

var (
	from = flag.Int("from", 0, "starting id to build the index from.")
	to   = flag.Int("to", 1, "ending index to build the index till.")
	term = flag.String("keyword", "", "search term or keyword for transcript and title")
)

func main() {
	log.SetLevel(log.InfoLevel)
	flag.Parse()
	index := xkcd.New(indexDB)

	if *from != 0 && *to != 0 {
		log.WithFields(log.Fields{
			"program":    "main.go",
			"index from": from,
			"index to":   to,
		}).Info("building index between the interval.")
		index.Build(*from, *to)
		index.Save()
	}

	if *term != "" {
		comics := index.Search(*term)
		re := regexp.MustCompile(`\[([^\[\]]*)\]`)
		for _, comic := range comics {
			fmt.Printf("%s\n", strings.Repeat("-", 75))
			fmt.Printf("Found Comic ID: %15v\n", comic.Num)
			fmt.Printf("URL: %15v\n", comic.Img)

			if comic.Transcript != "" {
				if re.MatchString(comic.Transcript) {
					fmt.Print("Transcript:\n")
					submatchall := re.FindAllString(comic.Transcript, -1)
					for _, element := range submatchall {
						element = strings.Trim(element, "[")
						element = strings.Trim(element, "]")
						fmt.Println(element)
					}
				} else {
					fmt.Printf("Transcript: %15v\n", comic.Transcript)
				}
			} else {
				if re.MatchString(comic.Alt) {
					fmt.Print("Alt:\n")
					submatchall := re.FindAllString(comic.Transcript, -1)
					for _, element := range submatchall {
						element = strings.Trim(element, "[")
						element = strings.Trim(element, "]")
						fmt.Println(element)
					}
				} else {
					fmt.Printf("Alt: %15v\n", comic.Alt)
				}
			}
			fmt.Printf("%s\n", strings.Repeat("-", 75))
		}
	}
}
