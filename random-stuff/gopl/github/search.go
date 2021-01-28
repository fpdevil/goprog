package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

// SearchIssues function queries the GitHub issue tracker.
func SearchIssues(queryparams []string) (*SearchResults, error) {
	q := url.QueryEscape(strings.Join(queryparams, " "))
	log.WithFields(log.Fields{
		"function": "SearchIssues",
	}).Infof("initiating issue search for %s\n", IssuesURL+"?q="+q)

	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		log.WithFields(
			log.Fields{
				"function": "SearchIssues",
				"error":    error.Error,
				"url":      IssuesURL + "?q=" + q,
			}).Error("error during GET")
	}

	// close the resp body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.WithFields(
			log.Fields{
				"function": "SearchIssues",
				"status":   resp.Status,
			}).Error("search query failed")
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	log.WithFields(
		log.Fields{
			"function": "SearchIssues",
			"status":   resp.Status,
		}).Infof("search query response status: %s", resp.Status)

	var result SearchResults
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		resp.Body.Close()
		log.WithFields(
			log.Fields{
				"function": "SearchIssues",
				"error":    err.Error(),
			}).Error("json decoding of results failed")
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
