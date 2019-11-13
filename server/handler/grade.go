package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

// Grade ...
func (h *handlerImpl) Grade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// read body as JSON
	var payload gradeReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
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

	// try to grade asynchronously
	if !h.grader.GradeAsync(
		newUUID, payload.SubmissionCode, payload.GradingCode, payload.GradingMethod, payload.MemoryLimitKB,
		payload.TimeLimitMS, payload.InputURL, payload.OutputURL) {
		h.logger.Warn("Grader is busy with status %s", h.grader.Status())
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&gradeRes{UUID: newUUID})
}
