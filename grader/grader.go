package grader

// Grader is the main component of this judge that will handle a submission at a time
//go:generate mockery -name=Grader
type Grader interface {
	// Status returns the current status of the grader
	Status() Status

	// TODO GradeAsync grades the submission asynchronously
}

type graderImpl struct {
	status Status
}

func (g *graderImpl) Status() Status {
	return g.status
}
