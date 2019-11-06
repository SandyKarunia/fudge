package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sandykarunia/fudge/grader"
	"net/http"
)

// Grade ...
func (h *handlerImpl) Grade(w http.ResponseWriter, r *http.Request) {
	// check if the judge is currently busy or not, if so, reject this request
	// use check-lock-check to be safe
	if h.grader.Status() != grader.StatusIdle {
		w.WriteHeader(http.StatusConflict)
		return
	}
	h.graderLock.Lock()
	defer h.graderLock.Unlock()
	if h.grader.Status() != grader.StatusIdle {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// read body as JSON
	var payload gradeReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		// TODO log error
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// validate the payload
	if errors := payload.validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string][]string{
			"errors": errors,
		})
		return
	}

	// create uuid for the grade request
	// this uuid will be used when the judge notifies the server
	newUUID := uuid.New().String()

	// TODO process the submission

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&gradeRes{UUID: newUUID})
}
