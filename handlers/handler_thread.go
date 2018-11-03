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
	"time"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateThread(w http.ResponseWriter, r *http.Request) {
	thread := models.Thread{}
	json.NewDecoder(r.Body).Decode(&thread)
	thread.Forum = mux.Vars(r)["slug"]

	if thread.Created == "" {
		t := time.Now()
		thread.Created = t.Format("2006-01-02T15:04:05.999999999Z07:00")
	}

	res, err := queries.ForumGetBySlug(h.DB, thread.Forum)
	if err == sql.ErrNoRows {
		errMsg := models.Error{}
		errMsg.ErrorForum(thread.Forum)
		network.ResponseNotFound(w, errMsg)
		return
	}
	thread.Forum = res.Slug

	err = queries.ThreadInsert(h.DB, &thread)
	if err != nil {
		log.Println("Can not create thread", err)

		res, err := queries.ThreadGetBySlug(h.DB, *thread.Slug)
		if err == sql.ErrNoRows {
			errMsg := models.Error{}
			errMsg.ErrorForum(thread.Forum)
			network.ResponseNotFound(w, errMsg)
			return
		}
		thread = *res

		network.ResponseConflict(w, thread)
		return
	}

	network.ResponseCreated(w, thread)
}

func (h *Handler) GetThreadsList(w http.ResponseWriter, r *http.Request) {
	thread := models.Thread{}
	thread.Forum = mux.Vars(r)["slug"]

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

	// threads := models.Threads{}
	// res := queries.ThreadQuery(h.DB, thread.Forum, limit, desc, since)
	// threads = *res
	threads := *queries.ThreadQuery(h.DB, thread.Forum, limit, desc, since)

	if len(threads) == 0 {
		res, err := queries.ForumGetBySlug(h.DB, thread.Forum)
		if res.Slug == "" || err != nil {
			errMsg := models.Error{}
			errMsg.ErrorForum(thread.Forum)
			network.ResponseNotFound(w, errMsg)
			return
		}
	}

	network.ResponseOK(w, threads)
}

func (h *Handler) GetThreadDetails(w http.ResponseWriter, r *http.Request) {
	thread := models.Thread{}
	slugORid := mux.Vars(r)["slug_or_id"]

	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		tmp, err := queries.ThreadGetBySlug(h.DB, slugORid)
		if err != nil {
			errMsg := models.Error{}
			errMsg.ErrorThreadBySlug(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		thread = *tmp
	} else {
		tmp, err := queries.ThreadGetByID(h.DB, threadID)
		if err != nil {
			errMsg := models.Error{}
			errMsg.ErrorThreadById(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		thread = *tmp
	}

	network.ResponseOK(w, thread)
}

func (h *Handler) UpdateThread(w http.ResponseWriter, r *http.Request) {
	thread := models.Thread{}
	json.NewDecoder(r.Body).Decode(&thread)
	slugORid := mux.Vars(r)["slug_or_id"]

	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		tmp, err := queries.ThreadGetBySlug(h.DB, slugORid)
		if err == sql.ErrNoRows {
			errMsg := models.Error{}
			errMsg.ErrorThreadBySlug(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		threadID = tmp.ID
	}
	thread.ID = threadID

	if _, err = queries.ThreadGetByID(h.DB, threadID); err == sql.ErrNoRows {
		errMsg := models.Error{}
		errMsg.ErrorThreadById(slugORid)
		network.ResponseNotFound(w, errMsg)
		return
	}

	err = queries.ThreadUpdate(h.DB, &thread)
	if err != nil {
		log.Println("Can not UPDATE", err)
		errMsg := models.Error{}
		errMsg.ErrorUser(*thread.Slug)
		network.ResponseConflict(w, errMsg)
		return
	}

	// _, err = queries.ThreadGetByID(h.DB, threadID)
	// switch {
	// case err == sql.ErrNoRows:
	// 	log.Println("Thread not found", err)
	// 	errMsg := models.Error{}
	// 	errMsg.ErrorUser(*thread.Slug)
	// 	network.ResponseNotFound(w, errMsg)
	// case err != nil:
	// 	log.Println("Thread doesn't exist", err)
	// 	return
	// default:
	// 	err = queries.ThreadUpdate(h.DB, &thread)
	// 	if err != nil {
	// 		log.Println("Can not UPDATE", err)
	// 		errMsg := models.Error{}
	// 		errMsg.ErrorUser(*thread.Slug)
	// 		network.ResponseConflict(w, errMsg)
	// 		return
	// 	}

	// }
	network.ResponseOK(w, thread)
}
