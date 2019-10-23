package handler

import (
	"encoding/json"
	"gitlab.com/sandykarunia/fudge/grader"
	"net/http"
)

// For test cases (input, output, inputURL, outputURL):
// - if input is supplied, output also has to be supplied
// - if inputURL is supplied, outputURL also has to be supplied
// - if input and inputURL are both supplied, input will be used and inputURL will be ignored
// - if none is supplied, then the request will be rejected
type payloadStr struct {
	// required, unique identifier for the request, it will be used in the webhook when judge is ready to return the result
	uid string `json:"uid"`

	// required, submission code to be compiled and graded
	submissionCode string `json:"submission_code"`

	// optional, can be GRADE_ALL or GRADE_UNTIL_FAIL, defaults to GRADE_UNTIL_FAIL
	gradingMethod grader.Method `json:"grading_method"`

	// required, memory limit ...
	memoryLimitKB int64 `json:"memory_limit_kb"`

	// required, time limit ...
	timeLimitMS int64 `json:"time_limit_ms"`

	// optional, usually is used if there are multiple solutions, for example: accepted when the output is odd number.
	// this code is required to output true / false value as stdout to decide whether the output of submissionCode is
	// correct or not. The judge will supply the output of submissionCode as stdin to this code.
	// If this is not supplied, judge will use simple grading method (compare string, ignore multiple whitespaces).
	// If this is supplied, judge will ignore "output" and "outputURL" in the payload, and use this code to judge
	// instead.
	gradingCode string `json:"grading_code"`

	// optional, list of input & output for test cases
	// if input and output have different lengths, the request will be rejected
	input  []string `json:"input"`
	output []string `json:"output"`

	// optional, list of input URL & output URL for test cases
	// if inputURL and outputURL have different ranges, the request will be rejected
	// inputURL / outputURL have to be in specific form:
	// - each of them has to have lengths == 2
	// - the first element is the URL, with '{%}' in it, example: "http://your-tc-url/my_tc-{%}.in"
	// - the second element is the range, which has to follow regex pattern '^\d+-\d+$', example: "1-5",
	//   if the range contains invalid range (e.g. "5-1"), then the request will be rejected.
	// the judge will iterate through the number range (e.g. "1-5" means iterate through "1", "2", ..., "5"),
	// and for each number, it will replace the URL in the first element.
	// For example, if the value is ["http://my-tc-url/tc_problem_a_{%}.in", "3-5"], then it will be expanded to:
	// - http://my-tc-url/tc_problem_a_3.in
	// - http://my-tc-url/tc_problem_a_4.in
	// - http://my-tc-url/tc_problem_a_5.in
	inputURL  []string `json:"input_url"`
	outputURL []string `json:"output_url"`
}

// Grade ...
func Grade(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload payloadStr
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
	panic("IMPLEMENT")
}
