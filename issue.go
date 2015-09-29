package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Issue is the internal representation of an issue
type Issue struct {
	ID       string    `json:"id"`
	Number   int       `json:"number"`
	Status   string    `json:"status"`
	Title    string    `json:"title,omitempty"`
	Body     string    `json:"body,omitempty"`
	Assignee string    `json:"assignee,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
	URL      string    `json:"url,omitempty"`
	Repo     string    `json:"repo"`
	Owner    string    `json:"owner"`
}

func (i *Issue) Path() string {
	return i.Owner + "/" + i.Repo + "/" + i.ID
}

// Comments is a collection/slice of comments
type Comments []Comment

// Comment is the internal representation of an issue comment
type Comment struct {
	ID        int       `json:"id,omitempty"`
	Body      string    `json:"body,omitempty"`
	User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	URL       string    `json:"url,omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	IssueURL  string    `json:"issue_url,omitempty"`
}

// Issues is a collection/slice of issues
type Issues []*Issue

type Error struct {
	Message string `json:"error"`
}

func printIssuesList(body []byte) error {
	issues := Issues{}
	err := json.Unmarshal(body, &issues)
	if err != nil {
		return err
	}

	for _, issue := range issues {
		title := color.YellowString(issue.Title)
		fmt.Printf("\n" + title + " - [" + issue.Path() + "]")
	}

	return nil
}

func printIssueDetails(body []byte) {
	issue := Issue{}
	err := json.Unmarshal(body, &issue)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Blue("[ #" + issue.ID + "] " + issue.Title)
	fmt.Println(issue.Body)
	fmt.Println(issue.Status)
}
