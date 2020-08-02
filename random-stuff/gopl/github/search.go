package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

// SearchIssues function queries the github issues tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	u := IssuesURL + "?q=" + q
	log.Infof("now calling the api %v\n", u)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	var result IssuesSearchResult
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return &result, nil
}
