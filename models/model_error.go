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

func (e *Error) ErrorParent() {
	e.Message = "Parent post was created in another thread"
}

func (e *Error) ErrorPostAuthor(author string) {
	e.Message = "Can't find post author by nickname: " + author
}

func (e *Error) ErrorPost(postID string) {
	e.Message = "Can't find post with id: " + postID
}

func (e *Error) ErrorPostThread(threadID string) {
	e.Message = "Can't find post thread by id: " + threadID
}

func (e *Error) ErrorThreadBySlug(threadSlug string) {
	e.Message = "Can't find thread by slug: " + threadSlug
}

func (e *Error) ErrorThreadById(threadID string) {
	e.Message = "Can't find thread by id: " + threadID
}
