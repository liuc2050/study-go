package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

//SearchIssues queries the Github issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failded: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

const issuelist = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//write
func SearchIssuesAndOutputByTemplate(terms []string) {
	result, err := SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	report, err := template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(issuelist)
	if err != nil {
		log.Fatal(err)
	}
	err = report.Execute(os.Stdout, result)
	if err != nil {
		log.Fatal(err)
	}
}
