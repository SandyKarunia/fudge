package language

type cpp struct{}

func (*cpp) GetVersionCmd() (string, []string) {
	return "g++", []string{"--version"}
}

func (*cpp) Name() string {
	return "CPP"
}
