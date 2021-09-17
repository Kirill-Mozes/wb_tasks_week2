package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simplerest/pkg/model"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}

	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !event.Validate() {
		http.Error(w, "wrong fields", http.StatusBadRequest)
		return
	}

	id, err := h.repos.CreateEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idBytes, err := json.Marshal(map[string]interface{}{
		"id": id,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(idBytes)
}
