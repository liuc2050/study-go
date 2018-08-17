package github

import (
	"fmt"
	"testing"
)

func TestSearchIssues(t *testing.T) {
	tests := [][]string{
		{"repo:golang/go", "is:open", "json", "decoder"},
	}

	for _, test := range tests {
		fmt.Printf("input:%v\n", test)
		result, err := SearchIssues(test)
		if err != nil {
			t.Errorf("error:%v", err)
			continue
		}
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func TestSearchIssuesAndOutputByTemplate(t *testing.T) {
	tests := [][]string{
		{"repo:golang/go", "is:open", "json", "decoder"},
	}

	for _, test := range tests {
		fmt.Printf("input:%v\n", test)
		SearchIssuesAndOutputByTemplate(test)
	}
}
