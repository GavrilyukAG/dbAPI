package handlers

import (
	"database/sql"

	"github.com/GavrilyukAG/dbAPI/models"
	"github.com/GavrilyukAG/dbAPI/network"
	"github.com/GavrilyukAG/dbAPI/queries"

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
		tmp, err := queries.ThreadGetBySlug(h.DB, slugORid) // slug
		if err != nil {
			errMsg := models.Error{}
			errMsg.ErrorThreadBySlug(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		thread = *tmp
	} else {
		tmp, err := queries.ThreadGetByID(h.DB, threadID) // ID
		if err != nil {
			errMsg := models.Error{}
			errMsg.ErrorThreadById(slugORid)
			network.ResponseNotFound(w, errMsg)
			return
		}
		thread = *tmp
	}

	if _, err = queries.UserGetByNickname(h.DB, vote.Nickname); err == sql.ErrNoRows {
		errMsg := models.Error{}
		errMsg.ErrorUser(vote.Nickname)
		network.ResponseNotFound(w, errMsg)
		return
	}

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
