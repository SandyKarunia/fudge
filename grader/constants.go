package grader

// Method is the grading method to be used by the grader
type Method string

// available grading methods
const (
	GradeAll       Method = "GRADE_ALL"
	GradeUntilFail        = "GRADE_UNTIL_FAIL"
)

// Status describes what the grader is currently doing
type Status string

// available status
const (
	StatusIdle         Status = "IDLE"
	StatusFetchInput          = "FETCH_INPUT"
	StatusFetchOutput         = "FETCH_OUTPUT"
	StatusPrepare             = "PREPARE"
	StatusGrading             = "GRADING"
	StatusNotifyResult        = "NOTIFY_RESULT"
)
