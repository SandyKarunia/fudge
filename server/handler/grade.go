package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/pure"
	"github.com/google/uuid"
	"github.com/sandykarunia/fudge/grader"
	"net/http"
	"regexp"
)

type reqPayloadStr struct {
	// required, submission code to be compiled and graded
	SubmissionCode string `json:"submission_code"`

	// optional, can be GRADE_ALL or GRADE_UNTIL_FAIL, defaults to GRADE_UNTIL_FAIL
	GradingMethod grader.Method `json:"grading_method"`

	// required, memory limit ...
	MemoryLimitKB int64 `json:"memory_limit_kb"`

	// required, time limit ...
	TimeLimitMS int64 `json:"time_limit_ms"`

	// optional, usually is used if there are multiple solutions, for example: accepted when the output is odd number.
	// this code is required to output true / false value as stdout to decide whether the output of submissionCode is
	// correct or not. The judge will supply the output of submissionCode as stdin to this code.
	// If this is not supplied, judge will use simple grading method (compare string, ignore multiple whitespaces).
	// If this is supplied, judge will ignore "output" and "outputURL" in the payload, and use this code to judge
	// instead.
	GradingCode string `json:"grading_code"`

	// required, list of input URL & output URL for test cases
	// if inputURL and outputURL have different ranges, the request will be rejected
	// inputURL / outputURL have to be in specific form:
	// - each of them has to have lengths == 2
	// - the first element is the URL, with '{%}' in it, example: "http://your-tc-url/my_tc-{%}.in"
	// - the second element is the range, which has to follow regex pattern '^\d+-\d+$', example: "1-5"
	// the judge will iterate through the number range (e.g. "1-5" means iterate through "1", "2", ..., "5"),
	// and for each number, it will replace the URL in the first element.
	// For example, if the value is ["http://my-tc-url/tc_problem_a_{%}.in", "3-5"], then it will be expanded to:
	// - http://my-tc-url/tc_problem_a_3.in
	// - http://my-tc-url/tc_problem_a_4.in
	// - http://my-tc-url/tc_problem_a_5.in
	InputURL  []string `json:"input_url"`
	OutputURL []string `json:"output_url"`
}

// Grade ...
func Grade(w http.ResponseWriter, r *http.Request) {
	// TODO check if the judge is currently judging or not, if so, reject this request

	w.Header().Set(pure.ContentType, pure.ApplicationJSON)

	// read body as JSON
	var payload reqPayloadStr
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
	var errors []string
	if len(payload.SubmissionCode) == 0 {
		errors = append(errors, "submission_code can't be empty")
	}
	if payload.GradingMethod != grader.GradeAll && payload.GradingMethod != grader.GradeUntilFail {
		errors = append(errors, fmt.Sprintf("grading_method must be either %s or %s", grader.GradeAll, grader.GradeUntilFail))
	}
	// hard code to 128MB
	maxMemoryLimitKB := int64(128 * 1024)
	if payload.MemoryLimitKB > maxMemoryLimitKB {
		errors = append(errors, fmt.Sprintf("memory_limit_kb value must be less or equal than %d", maxMemoryLimitKB))
	}
	// hard code to 1 minute
	maxTimeLimitMS := int64(60 * 1000)
	if payload.TimeLimitMS > maxTimeLimitMS {
		errors = append(errors, fmt.Sprintf("time_limit_ms value must be less or equal than %d", maxTimeLimitMS))
	}
	if len(payload.InputURL) != 2 {
		errors = append(errors, fmt.Sprintf("input_url must be an array with length equals to 2"))
	}
	if len(payload.OutputURL) != 2 {
		errors = append(errors, fmt.Sprintf("output_url must be an array with length equals to 2"))
	}
	if len(payload.InputURL) == 2 && len(payload.OutputURL) == 2 {
		rng := payload.InputURL[1]
		if rng != payload.OutputURL[1] {
			errors = append(errors, fmt.Sprintf("range (2nd element) in input_url has to be equal to output_url"))
		}
		if matched, _ := regexp.Match("^\\d+-\\d+$", []byte(rng)); !matched {
			errors = append(errors, fmt.Sprintf("range (2nd element) in input_url and output_url have to follow regex pattern '^\\d+-\\d+$', e.g. '3-51'"))
		}
	}

	if len(errors) > 0 {
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
	_ = json.NewEncoder(w).Encode(map[string]string{
		"uuid": newUUID,
	})
}
