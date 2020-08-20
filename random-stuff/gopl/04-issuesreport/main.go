package main

import (
	"fmt"
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

// daysAgo function calculates the number of days elapsed since issue surfaced
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// template creation
var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func noMust() {
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

func main() {
	noMust()
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := report.Execute(os.Stdout, result); err != nil {
	// 	log.Fatal(err)
	// }
}
