package handler

import (
	"encoding/json"
	"net/http"
)

// HealthCheck ...
func (*handlerImpl) HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]bool{
		"healthy": true,
	})
}
