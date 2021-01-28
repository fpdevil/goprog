package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/fpdevil/goprog/random-stuff/gopl/github"
)

// templ is a template describing the github issues to be reported
const (
	templ = `
<h1>{{.TotalCount}} issues:</h1>
<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'></a>{{.Title}}</td>
	</tr>
	{{end}}
</table>
`
)

func main() {
	tp, err := template.New("issueList").Parse(templ)
	issueList := template.Must(tp, err)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}
