package models

type Post struct {
	Author   string `json:"author,omitempty"`
	Created  string `json:"created,omitempty"`
	Forum    string `json:"forum,omitempty"`
	ID       int64  `json:"id,omitempty"`
	IsEdited bool   `json:"isedited,omitempty"`
	Message  string `json:"message,omitempty"`
	Parent   int64  `json:"parent,omitempty"`
	// Path     []byte `json:"-"`
	Thread int `json:"thread,omitempty"`
}

type PostFull struct {
	Author *User
	Forum  *Forum
	Post   *Post
	Thread *Thread
}

type PostUpdate struct {
	Message string
}

type Posts []*Post
