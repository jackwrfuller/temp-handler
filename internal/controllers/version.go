package controllers

import (
	"net/http"
)

func (h *BaseHandler) getVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Temp Handler dev"))
}
