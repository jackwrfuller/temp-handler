package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *BaseHandler) update(w http.ResponseWriter, r *http.Request) {
	var newData SensorData

	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	lock.Lock()
	data = newData
	lock.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated successfully\n")
}
