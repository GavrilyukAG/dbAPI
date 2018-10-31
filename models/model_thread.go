package models

type Thread struct {
	Author  string  `json:"author"`
	Created string  `json:"created, omitempty"`
	Forum   string  `json:"forum"`
	ID      int     `json:"id, omitempty"`
	Message string  `json:"message"`
	Slug    *string `json:"slug, omitempty"`
	Title   string  `json:"title"`
	Votes   int     `json:"votes, omitempty"`
}

type ThreadUpdate struct {
	Message string
	Title   string
}

type Threads []*Thread
