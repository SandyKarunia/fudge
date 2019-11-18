package language

// Python3 ...
type Python3 struct{}

// GetVersionCmd ...
func (*Python3) GetVersionCmd() (string, []string) {
	return "python3", []string{"--version"}
}
