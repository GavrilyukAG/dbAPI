package handlers

import (
	"database/sql"
	"dbAPI/models"
	"dbAPI/network"
	"dbAPI/queries"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Nickname = mux.Vars(r)["nickname"]

	_, err := queries.UserInsert(h.DB, user)
	if err != nil {
		var users models.Users
		tmp, _ := queries.UserGetAll(h.DB, user.Nickname, user.Email)
		// users = append(users, tmp)
		users = *tmp

		network.ResponseConflict(w, users)
		return
	}

	// ответ
	network.ResponseCreated(w, user)
}

func (h *Handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	user.Nickname = mux.Vars(r)["nickname"]

	tmp, err := queries.UserGetByNickname(h.DB, user.Nickname)

	switch {
	case err == sql.ErrNoRows:
		errMsg := models.Error{}
		errMsg.ErrorUser(user.Nickname)
		network.ResponseNotFound(w, errMsg)
	case err != nil:
		// log.Fatal(err)
		log.Println("USER doesn't exist", err)
		return
	default:
		user = *tmp
		network.ResponseOK(w, user)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Nickname = mux.Vars(r)["nickname"]

	_, err := queries.UserGetByNickname(h.DB, user.Nickname)

	switch {
	case err == sql.ErrNoRows:
		// log.Println("USER not found", err)
		errMsg := models.Error{}
		errMsg.ErrorUser(user.Nickname)
		network.ResponseNotFound(w, errMsg)
	case err != nil:
		// log.Fatal(err)
		log.Println("USER doesn't exist", err)
		return
	default:
		err = queries.UserUpdate(h.DB, &user)
		if err != nil {
			// log.Println("Cannot UPDATE", err)
			errMsg := models.Error{}
			errMsg.ErrorUser(user.Nickname)
			network.ResponseConflict(w, errMsg)
			return
		}

		network.ResponseOK(w, user)
	}
}

func (h *Handler) GetUsersList(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	var (
		limit int    = 0
		desc  bool   = false
		since string = ""
	)

	for k, _ := range r.URL.Query() {
		switch k {
		case "limit":
			limit, _ = strconv.Atoi(r.URL.Query().Get(k))
		case "desc":
			desc, _ = strconv.ParseBool(r.URL.Query().Get(k))
		case "since":
			since = r.URL.Query().Get(k)
		}
	}

	_, err := queries.ForumGetBySlug(h.DB, slug)
	if err == sql.ErrNoRows {
		errMsg := models.Error{}
		errMsg.ErrorForum(slug)
		network.ResponseNotFound(w, errMsg)
		return
	}

	var users models.Users
	tmp, err := queries.UserGetAllBySlug(h.DB, slug, since, limit, desc)
	if err != nil {
		log.Println(err)
		return
	}
	users = *tmp
	network.ResponseOK(w, users)

	// switch {
	// case err != nil:
	// 	network.ResponseNotFound(w, users)
	// 	return
	// case err == sql.ErrNoRows:
	// 	errMsg := models.Error{}
	// 	errMsg.ErrorForum(slug)
	// 	network.ResponseNotFound(w, errMsg)
	// 	return
	// default:
	// 	users = *tmp
	// 	network.ResponseOK(w, users)
	// 	return
	// }
}
