package language

// python3 ...
type python3 struct{}

// GetVersionCmd ...
func (*python3) GetVersionCmd() (string, []string) {
	return "python3", []string{"--version"}
}

func (*python3) Name() string {
	return "PYTHON3"
}
