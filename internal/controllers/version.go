package controllers

import (
	"encoding/json"
	"net/http"
)

type Version struct {
	Version string `json:"version"`
}

var version = Version{
		Version: "dev",
}

func (h *BaseHandler) getVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}
