package models

type Vote struct {
	Nickname string `json:"nickname"`
	Thread   int    `json:"-"`
	Voice    int    `json:"voice"`
}
