package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"simplerest/pkg/model"
)

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}

	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !event.Validate() {
		http.Error(w, "wrong fields", http.StatusBadRequest)
		return
	}

	updatedEvent, err := h.repos.UpdateEvent(event)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	eventBytes, _ := json.Marshal(updatedEvent)

	w.Header().Set("Content-Type", "application/json")
	w.Write(eventBytes)
}
