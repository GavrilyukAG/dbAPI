package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GavrilyukAG/dbAPI/models"
	"github.com/GavrilyukAG/dbAPI/network"
	"github.com/GavrilyukAG/dbAPI/queries"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateForum(w http.ResponseWriter, r *http.Request) {
	forum := models.Forum{}
	json.NewDecoder(r.Body).Decode(&forum)

	// проверяем есть ли данный пользователь
	_, err := queries.UserGetByNickname(h.DB, forum.User)
	if err != nil {
		// log.Println("User doesn't exist", err)
		errMsg := models.Error{}
		errMsg.ErrorUser(forum.User)
		network.ResponseNotFound(w, errMsg)
		return
	}

	// вставляем данные в таблицу форума
	err = queries.ForumInsert(h.DB, &forum)
	if err != nil {
		// log.Println(err)
		res, _ := queries.ForumGetByUsername(h.DB, forum.User)
		forum = *res
		network.ResponseConflict(w, forum)
		return
	}

	network.ResponseCreated(w, forum)
}

func (h *Handler) GetForumDetails(w http.ResponseWriter, r *http.Request) {
	forum := models.Forum{}
	forum.Slug = mux.Vars(r)["slug"]

	res, err := queries.ForumGetBySlug(h.DB, forum.Slug)
	if err != nil {
		// log.Println(err)
		errMsg := models.Error{}
		errMsg.ErrorForum(forum.Slug)
		network.ResponseNotFound(w, errMsg)
		return
	}
	forum = *res

	network.ResponseOK(w, forum)
}
