package language

// Cpp ...
type Cpp struct{}

// GetVersionCmd ...
func (*Cpp) GetVersionCmd() (string, []string) {
	return "g++", []string{"--version"}
}
