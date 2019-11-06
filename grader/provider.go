package grader

// Provider ...
func Provider() Grader {
	obj := &graderImpl{}
	obj.status = StatusIdle

	return obj
}
