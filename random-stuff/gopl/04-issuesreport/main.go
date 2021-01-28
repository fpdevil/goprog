package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/fpdevil/goprog/random-stuff/gopl/github"
)

// templ is a template describing the github issues to be reported
const (
	templ = `{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`
)

//!+daysAgo
// daysAgo function calculates the number of days elapsed since the
// issue has surfaced
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//!-daysAgo

//!+noMust
// noMust function
func noMust() {
	// create a new template with name `report`
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Printf("error parsing result: %v\n", err)
		return
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}

//!-noMust

//+!execution
// template creation
var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	// noMust()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//-!execution
