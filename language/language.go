package language

// Language is an interface for a programming language supported in fudge
type Language interface {
	// Name returns the name of the programming language
	Name() string

	// VersionCmd returns the cmd and arguments to check the version of the programming language in current machine
	VersionCmd() (string, []string)

	// CompileCmd returns the cmd and arguments to compile a file in this language
	CompileCmd(filename string, outputFilename string) (string, []string)
}

// list of all languages available, since they are stateless, we can store them globally
var (
	CPP     = &cpp{}
	Python3 = &python3{}
)

var allLanguages = []Language{CPP, Python3}

// Get is a helper function to get a language from language name,
// if language with given name not found, then return nil
var Get = func(name string) Language {
	for _, l := range allLanguages {
		if name == l.Name() {
			return l
		}
	}
	return nil
}
