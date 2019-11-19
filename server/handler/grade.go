package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/language"
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

	// build grade async payload
	gradeAsyncPayload := &grader.GradeAsyncPayload{
		UUID:               newUUID,
		SubmissionCode:     payload.SubmissionCode,
		SubmissionLanguage: language.Get(payload.SubmissionLanguage),
		GradingCode:        payload.GradingCode,
		GradingLanguage:    language.Get(payload.GradingLanguage),
		GradingMethod:      payload.GradingMethod,
		MemoryLimitKB:      payload.MemoryLimitKB,
		TimeLimitMS:        payload.TimeLimitMS,
		InputURL:           payload.InputURL,
		OutputURL:          payload.OutputURL,
	}

	// try to grade asynchronously
	if !h.grader.GradeAsync(gradeAsyncPayload) {
		h.logger.Warn("Grader is busy with status %s", h.grader.Status())
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&gradeRes{UUID: newUUID})
}
