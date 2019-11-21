package handler

import (
	"fmt"
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/language"
	"regexp"
	"strings"
)

type gradeReq struct {
	// required, submission code to be compiled and graded
	SubmissionCode string `json:"submission_code"`

	// required, programming language of the submission, language has to be valid
	SubmissionLanguage string `json:"submission_language"`

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

	// optional, required if GradingCode exists
	GradingLanguage string `json:"grading_language"`

	// required, list of input URL & output URL for test cases
	// if inputURLPatternAndRange and outputURLPatternAndRange have different ranges, the request will be rejected
	// inputURLPatternAndRange / outputURLPatternAndRange have to be in specific form:
	// - each of them has to have lengths == 2
	// - the first element is the URL, with '{%}' in it, example: "http://your-tc-url/my_tc-{%}.in"
	// - the second element is the range, which has to follow regex pattern '^\d+-\d+$', example: "1-5"
	// the judge will iterate through the number range (e.g. "1-5" means iterate through "1", "2", ..., "5"),
	// and for each number, it will replace the URL in the first element.
	// For example, if the value is ["http://my-tc-url/tc_problem_a_{%}.in", "3-5"], then it will be expanded to:
	// - http://my-tc-url/tc_problem_a_3.in
	// - http://my-tc-url/tc_problem_a_4.in
	// - http://my-tc-url/tc_problem_a_5.in
	InputURLPatternAndRange  []string `json:"input_url_pattern_and_range"`
	OutputURLPatternAndRange []string `json:"output_url_pattern_and_range"`
}

func (payload *gradeReq) validate() []string {
	var errors []string

	// check submission code
	if len(payload.SubmissionCode) == 0 {
		errors = append(errors, "submission_code can't be empty")
	}

	// check submission language
	if language.Get(payload.SubmissionLanguage) == nil {
		errors = append(errors, fmt.Sprintf("unknown submission language: %s", payload.SubmissionLanguage))
	}

	// check grading language if grading code is not empty
	if len(payload.GradingCode) > 0 && language.Get(payload.GradingLanguage) == nil {
		errors = append(errors, fmt.Sprintf("unknown grading language: %s", payload.GradingLanguage))
	}

	// check grading method
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

	// check input_url_pattern_and_range length
	if len(payload.InputURLPatternAndRange) != 2 {
		errors = append(errors, fmt.Sprintf("input_url_pattern_and_range must be an array with length equals to 2"))
	}
	// check if input_url_pattern_and_range contains mandatory "{%}" string
	if len(payload.InputURLPatternAndRange) > 0 && !strings.Contains(payload.InputURLPatternAndRange[0], "{%}") {
		errors = append(errors, "URL string (1st element) in input_url_pattern_and_range have to contain '{%}' string")
	}

	// check output_url_pattern_and_range only if grading_code is empty
	if len(payload.GradingCode) == 0 {
		// check output_url_pattern_and_range length
		if len(payload.OutputURLPatternAndRange) != 2 {
			errors = append(errors, fmt.Sprintf("output_url_pattern_and_range must be an array with length equals to 2"))
		}
		// check if output_url_pattern_and_range contains mandatory "{%}" string
		if len(payload.OutputURLPatternAndRange) > 0 && !strings.Contains(payload.OutputURLPatternAndRange[0], "{%}") {
			errors = append(errors, "URL string (1st element) in output_url_pattern_and_range have to contain '{%}' string")
		}
		// check relation between input_url_pattern_and_range and output_url_pattern_and_range
		if len(payload.InputURLPatternAndRange) == 2 && len(payload.OutputURLPatternAndRange) == 2 {
			// check if input_url_pattern_and_range and output_url_pattern_and_range have same ranges
			rng := payload.InputURLPatternAndRange[1]
			if rng != payload.OutputURLPatternAndRange[1] {
				errors = append(errors, fmt.Sprintf("range (2nd element) in input_url_pattern_and_range has to be equal to output_url_pattern_and_range"))
			}
			// check if the range is valid
			if matched, _ := regexp.Match("^\\d+-\\d+$", []byte(rng)); !matched {
				errors = append(errors, fmt.Sprintf("range (2nd element) in input_url_pattern_and_range and output_url_pattern_and_range have to follow regex pattern '^\\d+-\\d+$', e.g. '3-51'"))
			}
		}
	}

	return errors
}

type gradeRes struct {
	UUID string `json:"uuid"`
}
