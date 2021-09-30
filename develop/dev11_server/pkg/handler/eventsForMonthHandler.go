package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}

	today := time.Now().Format("2006-01-02")
	afterMonth := time.Now().AddDate(0, 1, 0).Format("2006-01-02")

	events, err := h.repos.GetEventsForCountDays(today, afterMonth)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	eventsBytes, _ := json.Marshal(events)

	w.Header().Set("Content-Type", "application/json")
	w.Write(eventsBytes)
}
