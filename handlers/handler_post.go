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
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	posts := models.Posts{}
	json.NewDecoder(r.Body).Decode(&posts)
	slugORid := mux.Vars(r)["slug_or_id"]

	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		res, err := queries.ThreadGetBySlug(h.DB, slugORid)
		if err != nil {
			// log.Println("Such thread does not exist", err)
			errMsg := models.Error{}
			errMsg.ErrorThreadBySlug(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		threadID = res.ID
	}

	if _, err = queries.ThreadGetByID(h.DB, threadID); err == sql.ErrNoRows {
		id := strconv.Itoa(threadID)
		errMsg := models.Error{}
		errMsg.ErrorPostThread(id)
		network.ResponseNotFound(w, errMsg)
		return
	}

	t := time.Now()
	created := t.Format("2006-01-02T15:04:05.999999999Z07:00")

	for _, post := range posts {
		if post.Created == "" {
			post.Created = created
		}
		if post.Forum == "" {
			res, err := queries.ThreadGetByID(h.DB, threadID)
			if err != nil {
				log.Println("Such forum does not exist", err)
			}
			post.Forum = res.Forum
		}

		post.Thread = threadID

		if post.Parent != 0 {
			parent, err := queries.PostGetById(h.DB, int(post.Parent))
			if err != nil || post.Thread != parent.Thread {
				errMsg := models.Error{}
				errMsg.ErrorParent()
				network.ResponseConflict(w, errMsg)
				return
			}
		}

		if _, err = queries.UserGetByNickname(h.DB, post.Author); err == sql.ErrNoRows {
			errMsg := models.Error{}
			errMsg.ErrorPostAuthor(post.Author)
			network.ResponseNotFound(w, errMsg)
			return
		}

		err = queries.PostInsert(h.DB, post)
		if err != nil {
			log.Println("Cannot create post", err)
		}
	}

	network.ResponseCreated(w, posts)
}

func (h *Handler) GetThreadPosts(w http.ResponseWriter, r *http.Request) {
	slugORid := mux.Vars(r)["slug_or_id"]

	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		res, err := queries.ThreadGetBySlug(h.DB, slugORid)
		if err != nil {
			errMsg := models.Error{}
			errMsg.ErrorThreadBySlug(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		threadID = res.ID
	} else {
		_, err := queries.ThreadGetByID(h.DB, threadID)
		if err != nil {
			errMsg := models.Error{}
			errMsg.ErrorThreadById(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
	}

	var (
		limit int    = 0
		desc  bool   = false
		since int    = 0
		sort  string = ""
	)

	for k, _ := range r.URL.Query() {
		switch k {
		case "limit":
			limit, _ = strconv.Atoi(r.URL.Query().Get(k))
		case "desc":
			desc, _ = strconv.ParseBool(r.URL.Query().Get(k))
		case "since":
			since, _ = strconv.Atoi(r.URL.Query().Get(k))
		case "sort":
			sort = r.URL.Query().Get(k)
		}
	}

	posts := models.Posts{}
	switch sort {
	case "tree":
		posts = *queries.PostGetTree(h.DB, threadID, limit, desc, since)
	case "parent_tree":
		posts = *queries.PostGetParentTree(h.DB, threadID, limit, desc, since)
	default:
		posts = *queries.PostGetFlat(h.DB, threadID, limit, desc, since)
	}
	network.ResponseOK(w, posts)
}

func (h *Handler) GetPostDetails(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	postID, _ := strconv.Atoi(id)

	vars := r.URL.Query().Get("related")
	related := strings.Split(vars, ",")

	if _, err := queries.PostGetById(h.DB, postID); err == sql.ErrNoRows {
		errMsg := models.Error{}
		errMsg.ErrorPost(id)
		network.ResponseNotFound(w, errMsg)
		return
	}

	postDetails := models.PostFull{}
	tmp, err := queries.PostGetDetails(h.DB, postID, related)
	if err != nil {
		log.Println(err)
	}
	postDetails = *tmp

	network.ResponseOK(w, postDetails)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	postID, _ := strconv.Atoi(id)

	postMsg := models.PostUpdate{}           // var message string
	json.NewDecoder(r.Body).Decode(&postMsg) // messgae

	post := models.Post{}
	if res, err := queries.PostGetById(h.DB, postID); err != nil {
		errMsg := models.Error{}
		errMsg.ErrorPost(id)
		network.ResponseNotFound(w, errMsg)
		return
	} else {
		post = *res
	}

	if postMsg.Message != "" {
		tmp, err := queries.PostUpdate(h.DB, postID, postMsg.Message)
		if err != nil {
			log.Println("Cannot update post", err)
		}
		post = *tmp
	}

	network.ResponseOK(w, post)
}
