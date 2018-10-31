package network

import (
	"encoding/json"
	"net/http"
)

func ResponseOK(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	msg, _ := json.Marshal(i)
	w.Write(msg)
}

func ResponseCreated(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	msg, _ := json.Marshal(i)
	w.Write(msg)
}

func ResponseConflict(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusConflict)
	msg, _ := json.Marshal(i)
	w.Write(msg)
}

func ResponseNotFound(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	msg, _ := json.Marshal(i)
	w.Write(msg)
}
