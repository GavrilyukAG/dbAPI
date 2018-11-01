package handlers

import (
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

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	posts := models.Posts{}
	json.NewDecoder(r.Body).Decode(&posts)
	slugORid := mux.Vars(r)["slug_or_id"]

	// var threadID int
	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		res, err := queries.ThreadGetBySlug(h.DB, slugORid)
		if err != nil {
			log.Println("Such thread does not exist", err)
		}
		threadID = res.ID
	} else {
		_, err := queries.ThreadGetByID(h.DB, threadID)
		if err != nil {
			log.Println("Such thread does not exist", err)
		}
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
		err = queries.PostInsert(h.DB, post)
		if err != nil {
			log.Println("Can not create post", err)
		}
		// log.Println(post.Path)
		// if post.Parent != 0 {
		// 	path := []int64{}
		// 	// path = GetPostParentPath
		// 	path = append(path, post.ID)
		// 	// UpdatePostPath
		// } else {
		// 	// path := []int64{post.ID}
		// 	post.Path = append(post.Path, post.ID)
		// }
	}

	network.ResponseCreated(w, posts)
}

func (h *Handler) GetThreadPosts(w http.ResponseWriter, r *http.Request) {
	slugORid := mux.Vars(r)["slug_or_id"]

	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		res, err := queries.ThreadGetBySlug(h.DB, slugORid)
		if err != nil {
			log.Println("Such thread does not exist", err)
		}
		threadID = res.ID
	} else {
		_, err := queries.ThreadGetByID(h.DB, threadID)
		if err != nil {
			log.Println("Such thread does not exist", err)
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
