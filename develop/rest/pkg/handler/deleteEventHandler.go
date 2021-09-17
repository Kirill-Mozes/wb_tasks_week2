package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}

	type IdEvent struct {
		Id int64 `json:"id"`
	}

	var idStruct IdEvent
	err := json.NewDecoder(r.Body).Decode(&idStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.repos.DeleteEvent(idStruct.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
