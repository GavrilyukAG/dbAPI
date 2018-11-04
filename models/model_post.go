package models

type Post struct {
	Author   string `json:"author, omitempty"`
	Created  string `json:"created, omitempty"`
	Forum    string `json:"forum, omitempty"`
	ID       int64  `json:"id, omitempty"`
	IsEdited bool   `json:"isEdited, omitempty"`
	Message  string `json:"message, omitempty"`
	Parent   int64  `json:"parent, omitempty"`
	Thread   int    `json:"thread,omitempty"`
}

type PostFull struct {
	Author *User   `json:"author, omitempty"`
	Forum  *Forum  `json:"forum, omitempty"`
	Post   *Post   `json:"post, omitempty"`
	Thread *Thread `json:"thread, omitempty"`
}

type PostUpdate struct {
	Message string
}

type Posts []*Post
