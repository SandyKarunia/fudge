package language

type cpp struct{}

func (*cpp) VersionCmd() (string, []string) {
	return "g++", []string{"--version"}
}

func (*cpp) Name() string {
	return "CPP"
}

func (*cpp) CompileCmd(filename string, outputFilename string) (string, []string) {
	return "g++", []string{filename, "-o", outputFilename}
}
