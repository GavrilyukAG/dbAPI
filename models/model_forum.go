package models

type Forum struct {
	Posts   int64  `json:"posts, omitempty"`
	Slug    string `json:"slug"`
	Threads int    `json:"threads"`
	Title   string `json:"title"`
	User    string `json:"user"`
}
