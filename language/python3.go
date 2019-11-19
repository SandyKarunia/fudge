package language

// python3 ...
type python3 struct{}

// VersionCmd ...
func (*python3) VersionCmd() (string, []string) {
	return "python3", []string{"--version"}
}

func (*python3) Name() string {
	return "PYTHON3"
}

func (*python3) CompileCmd(filename string, outputFilename string) (string, []string) {
	return "mv", []string{filename, outputFilename}
}
