// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Linkable interface {
	IsLinkable()
}

type Searchable interface {
	IsSearchable()
}

type AddComment struct {
	Content string `json:"content"`
	PostID  string `json:"post_id"`
}

type EditBook struct {
	ID          *string `json:"id"`
	Title       *string `json:"title"`
	GoodreadsID string  `json:"goodreads_id"`
}

type EditPage struct {
	ID       *string `json:"id"`
	Slug     *string `json:"slug"`
	Content  string  `json:"content"`
	Title    string  `json:"title"`
	Category *string `json:"category"`
}

type EditPost struct {
	ID       *string    `json:"id"`
	Content  *string    `json:"content"`
	Title    *string    `json:"title"`
	Datetime *time.Time `json:"datetime"`
	Draft    *bool      `json:"draft"`
}

type Limit struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type NewGeo struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type NewLink struct {
	Title       string     `json:"title"`
	URI         URI        `json:"uri"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Created     *time.Time `json:"created"`
}

type NewLog struct {
	Code        string  `json:"code"`
	Description *string `json:"description"`
	Location    *NewGeo `json:"location"`
	Project     string  `json:"project"`
	Duration    *string `json:"duration"`
}

type NewStat struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type NewTweet struct {
	FavoriteCount int       `json:"favorite_count"`
	Hashtags      []string  `json:"hashtags"`
	ID            string    `json:"id"`
	Posted        time.Time `json:"posted"`
	RetweetCount  int       `json:"retweet_count"`
	Symbols       []string  `json:"symbols"`
	Text          string    `json:"text"`
	Urls          []*URI    `json:"urls"`
	ScreenName    string    `json:"screen_name"`
	UserMentions  []string  `json:"user_mentions"`
}

// A stat is a key value pair of two interesting strings.
type Stat struct {
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	Modified time.Time `json:"modified"`
}

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleNormal Role = "normal"
)

var AllRole = []Role{
	RoleAdmin,
	RoleNormal,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleNormal:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
