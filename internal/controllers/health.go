package controllers

import (
	"net/http"
)

func (h *BaseHandler) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
