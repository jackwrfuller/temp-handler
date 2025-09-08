package controllers

import (
	"encoding/json"
	"net/http"
)


func (h *BaseHandler) getStatus(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
