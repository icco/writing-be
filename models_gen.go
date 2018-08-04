// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package writing_be

import (
	time "time"
)

type Comment struct {
	ID string `json:"Id"`
}
type Link struct {
	ID          string    `json:"Id"`
	Title       string    `json:"Title"`
	URI         string    `json:"Uri"`
	Created     time.Time `json:"Created"`
	Description string    `json:"Description"`
	Screenshot  string    `json:"Screenshot"`
	Tags        []*string `json:"Tags"`
}
type NewLink struct {
	Title       string    `json:"Title"`
	URI         string    `json:"Uri"`
	Description string    `json:"Description"`
	Tags        []*string `json:"Tags"`
	Created     time.Time `json:"Created"`
}
type NewPost struct {
	Content string `json:"Content"`
	Title   string `json:"Title"`
}
type Post struct {
	ID          string    `json:"Id"`
	Title       string    `json:"Title"`
	Content     string    `json:"Content"`
	HTML        string    `json:"Html"`
	SummaryHTML string    `json:"SummaryHtml"`
	Readtime    int       `json:"Readtime"`
	Datetime    time.Time `json:"Datetime"`
	Created     time.Time `json:"Created"`
	Modified    time.Time `json:"Modified"`
	Draft       bool      `json:"Draft"`
	Tags        []*string `json:"Tags"`
}
