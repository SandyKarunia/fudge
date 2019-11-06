package grader

// Provider ...
func Provider() Grader {
	return &graderImpl{}
}
