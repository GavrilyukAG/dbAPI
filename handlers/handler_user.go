package handlers

import (
	"database/sql"
	"dbAPI/models"
	"dbAPI/network"
	"dbAPI/queries"
	"encoding/json"
	"log"
	"net/http"

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
			// log.Println("Can not UPDATE", err)
			errMsg := models.Error{}
			errMsg.ErrorUser(user.Nickname)
			network.ResponseConflict(w, errMsg)
			return
		}

		network.ResponseOK(w, user)
	}
}
