package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fpdevil/goprog/random-stuff/gopl/github"
	log "github.com/sirupsen/logrus"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.WithFields(log.Fields{
			"function": "main",
		}).Error("not enough arguments submitted.")
		fmt.Fprintf(os.Stderr, "usage: %s <repo_name query_params...>\n", filepath.Base(args[0]))
		return
	}
	result, err := github.SearchIssues(args[1:])
	if err != nil {
		log.WithFields(log.Fields{
			"function": "main",
			"error":    err.Error(),
		}).Error("error occurred while searching for issues.")
		return
	}

	fmt.Fprintf(os.Stdout, "%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Fprintf(os.Stdout, "#%5d %10.10s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
