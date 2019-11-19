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
	StatusAcknowledged          Status = "ACKNOWLEDGED"
	StatusCleanUp                      = "CLEAN_UP"
	StatusCompiling                    = "COMPILING"
	StatusFetchInput                   = "FETCH_INPUT"
	StatusFetchOutput                  = "FETCH_OUTPUT"
	StatusGrading                      = "GRADING"
	StatusIdle                         = "IDLE"
	StatusNotifyResult                 = "NOTIFY_RESULT"
	StatusPrepareSandbox               = "PREPARE_SANDBOX"
	StatusPrepareSubmissionCode        = "PREPARE_SUBMISSION_CODE"
)
