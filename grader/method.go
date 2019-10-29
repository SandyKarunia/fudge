package grader

// Method ...
type Method string

// available grading methods
const (
	GradeAll       Method = "GRADE_ALL"
	GradeUntilFail        = "GRADE_UNTIL_FAIL"
)
