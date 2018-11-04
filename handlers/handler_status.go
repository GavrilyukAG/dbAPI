package handlers

import (
	"github.com/GavrilyukAG/dbAPI/models"
	"github.com/GavrilyukAG/dbAPI/network"
	"github.com/GavrilyukAG/dbAPI/queries"

	"log"
	"net/http"
)

func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	status := models.Status{}
	tmp, err := queries.StatusGet(h.DB)
	if err != nil {
		log.Println(err)
	}
	status = *tmp

	network.ResponseOK(w, status)
}

func (h *Handler) EraseDB(w http.ResponseWriter, r *http.Request) {
	err := queries.StatusClear(h.DB)
	if err != nil {
		log.Println(err)
	}
	network.ResponseOK(w, nil)
}
