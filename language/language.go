package language

// Language is an interface for a programming language supported in fudge
type Language interface {
	// GetVersionCmd returns the cmd and arguments to check the version of the programming language in current machine
	GetVersionCmd() (string, []string)
}
