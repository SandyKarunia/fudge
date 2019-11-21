package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/language"
	"net/http"
	"strconv"
	"strings"
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

	// split inputURLPatternAndRange and outputURLPatternAndRange to array of strings, replace the {%} pattern
	var inputURLs, outputURLs []string
	inputURLs = splitInputOutputPattern(payload.InputURLPatternAndRange)
	// split outputURLPatternAndRange only if gradingCode is empty
	if len(payload.GradingCode) == 0 {
		outputURLs = splitInputOutputPattern(payload.OutputURLPatternAndRange)
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
		InputURL:           inputURLs,
		OutputURL:          outputURLs,
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

// helper function to split the input / output from pattern
func splitInputOutputPattern(patternAndRange []string) []string {
	split := strings.Split(patternAndRange[1], "-")
	from, _ := strconv.Atoi(split[0])
	to, _ := strconv.Atoi(split[1])
	if from > to {
		temp := from
		from = to
		to = temp
	}

	var res []string
	for i := from; i <= to; i++ {
		res = append(res, strings.ReplaceAll(patternAndRange[0], "{%}", strconv.Itoa(i)))
	}

	return res
}
