package main

import (
	"fmt"
	"os"

	"github.com/fpdevil/goprog/random-stuff/gopl/github"
	log "github.com/sirupsen/logrus"
)

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Errorf("%v\n", err)
		return
	}

	fmt.Printf("%d issues:\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
