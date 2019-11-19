package sniffers

import (
	"github.com/fatih/color"
	"github.com/sandykarunia/fudge/language"
	"github.com/sandykarunia/fudge/utils"
)

type message struct {
	success string
	fail    string
}

// Sniffers contains the functional logic for groundcheck to use
// unlike Checkers, all the function in sniffers don't return anything, as they are optionals
//go:generate mockery -name=Sniffers
type Sniffers interface {
	SniffControlGroupSupport()
	SniffLanguageCppSupport()
	SniffLanguagePython3Support()
}

type sniffersImpl struct {
	sysUtils utils.System
}

func (s *sniffersImpl) SniffControlGroupSupport() {
	msg := &message{
		success: "Control group is supported",
		fail:    "Control group is not supported, sandbox won't allow program to start multiple processes of threads",
	}
	ok := s.sysUtils.IsControlGroupSupported()
	printPretty(ok, msg)
}

func (s *sniffersImpl) SniffLanguageCppSupport() {
	msg := &message{
		success: "SubmissionLanguage: C++ is supported",
		fail:    "SubmissionLanguage: C++ is not supported",
	}
	printPretty(s.sysUtils.IsLanguageSupported(language.CPP), msg)
}

func (s *sniffersImpl) SniffLanguagePython3Support() {
	msg := &message{
		success: "SubmissionLanguage: Python 3 is supported",
		fail:    "SubmissionLanguage: Python 3 is not supported",
	}
	printPretty(s.sysUtils.IsLanguageSupported(language.Python3), msg)
}

func printPretty(ok bool, msg *message) {
	if ok {
		color.Green("[ok] %s", msg.success)
		return
	}

	color.Yellow("[!!] %s", msg.fail)
}
