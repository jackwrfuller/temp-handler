package controllers 

import (
	"net/http"
	"sync"
)

type BaseHandler struct {

}

func NewBaseHandler() (*BaseHandler) {
	return &BaseHandler{}
}

type SensorData struct {
	Temp float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
}

var (
	data SensorData
	lock sync.Mutex
)

func (h *BaseHandler) HandleRequests(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/api/v1/version" && req.Method == http.MethodGet {
		h.getVersion(w, req)
		return
	}

	if req.URL.Path == "/api/v1/health" && req.Method == http.MethodGet {
		h.healthcheck(w, req)
		return
	}

	if req.URL.Path == "/api/v1/status" && req.Method == http.MethodGet {
		h.getStatus(w, req)
		return
	}

	if req.URL.Path == "/api/v1/update" && req.Method == http.MethodPost {
		h.update(w, req)
		return
	}

	http.NotFound(w, req)
	return
}
