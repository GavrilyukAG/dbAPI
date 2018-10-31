package models

type Error struct {
	Message string `json:"message"`
}

func (e *Error) ErrorUser(userNick string) {
	e.Message = "Can't find user with nickname " + userNick
}

func (e *Error) ErrorForum(forumSlug string) {
	e.Message = "Can't find forum with slug " + forumSlug
}
