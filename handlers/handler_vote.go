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
)

func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	json.NewDecoder(r.Body).Decode(&vote)
	slugORid := mux.Vars(r)["slug_or_id"]

	thread := models.Thread{}
	threadID, err := strconv.Atoi(slugORid)
	if err != nil {
		tmp, _ := queries.ThreadGetBySlug(h.DB, slugORid) // slug
		thread = *tmp
	} else {
		tmp, _ := queries.ThreadGetByID(h.DB, threadID) // ID
		thread = *tmp
	}

	// if err != nil {
	// 	log.Println("Such thread does not exist", err)
	// 	return
	// }

	vote.Thread = thread.ID
	oldVote, err := queries.VoteGetByNickname(h.DB, vote.Nickname)
	switch {
	case err == sql.ErrNoRows:
		_, err = queries.VoteInsert(h.DB, vote)
		if err != nil {
			log.Println(err)
			return
		}
		switch vote.Voice {
		case 1:
			_ = queries.ThreadVoteIncr(h.DB, &thread)
		case -1:
			_ = queries.ThreadVoteDecr(h.DB, &thread)
		}
	case err != nil:
		log.Println(err)
	default:
		if oldVote.Voice == vote.Voice {
			log.Println("Vote already exisists")
		} else {
			switch oldVote.Voice {
			case 1:
				_ = queries.ThreadVoteDecr(h.DB, &thread)
				_ = queries.ThreadVoteDecr(h.DB, &thread)
			case -1:
				_ = queries.ThreadVoteIncr(h.DB, &thread)
				_ = queries.ThreadVoteIncr(h.DB, &thread)
			}
			_ = queries.VoteUpdate(h.DB, &vote)
		}
	}

	network.ResponseOK(w, thread)
}
