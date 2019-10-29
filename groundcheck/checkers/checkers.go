package checkers

import (
	"github.com/fatih/color"
	"github.com/sandykarunia/fudge/utils"
)

type message struct {
	success string
	fail    string
	solve   []string
}

// Checkers contains the functional logic for groundcheck to use
//go:generate mockery -name=Checkers
type Checkers interface {
	// CheckSudo Checks whether we have root privilege or not
	CheckSudo() bool

	// CheckLibcapDevPkg Checks whether libcap-dev package is installed or not
	CheckLibcapDevPkg() bool

	// CheckIsolateBinaryExists Checks whether isolate binary exists or not
	CheckIsolateBinaryExists() bool

	// CheckIsolateBinaryExecutable Checks whether isolate binary is real or not
	// we do this by checking the version of the isolate binary
	CheckIsolateBinaryExecutable() bool
}

type checkersImpl struct {
	sysUtils  utils.System
	fileUtils utils.File
	pathUtils utils.Path
}

func (c *checkersImpl) CheckSudo() bool {
	msg := &message{
		success: "Program is running with root privilege",
		fail:    "Program is running without root privilege, we need root to run the isolate (sandbox) binary",
		solve:   []string{"run the binary with sudo"},
	}

	isSudo := c.sysUtils.IsSudo()
	printPretty(isSudo, msg)

	return isSudo
}

func (c *checkersImpl) CheckLibcapDevPkg() bool {
	msg := &message{
		success: "Required libcap-dev package is installed",
		fail:    "Required libcap-dev package is missing",
		solve: []string{
			"install libcap-dev package with package manager (e.g. \"sudo apt-get libcap-dev\")",
			"verify installation: dpkg -s libcap-dev",
		},
	}

	err := c.sysUtils.VerifyPkgInstalled("libcap-dev")
	if err != nil {
		// TODO log error
	}
	printPretty(err == nil, msg)
	return err == nil
}

func (c *checkersImpl) CheckIsolateBinaryExists() bool {
	isolatePath := c.pathUtils.IsolateBinary()
	fudgeDir := c.sysUtils.GetFudgeDir()
	msg := &message{
		success: "Required isolate binary found in " + isolatePath,
		fail:    "Required isolate binary not found in " + isolatePath,
		solve: []string{
			"go to github.com/ioi/isolate/releases",
			"click on the latest release tag",
			"download the source code",
			"extract the source code anywhere you want",
			"inside the extracted folder, run \"make isolate\" in command line, this requires libcap-dev library",
			"there should be a generated binary \"isolate\"",
			"copy \"isolate\" binary into " + fudgeDir + " directory",
		},
	}
	exists := c.fileUtils.Exists(isolatePath)
	printPretty(exists, msg)
	return exists
}

func (c *checkersImpl) CheckIsolateBinaryExecutable() bool {
	isolatePath := c.pathUtils.IsolateBinary()
	msg := &message{
		success: "Required isolate binary is executable",
		fail:    "Required isolate binary is not executable",
		solve: []string{
			"make sure the isolate binary is executable by running '" + isolatePath + " --version' in your command line",
			"if you got permission denied error, try running 'chmod +x " + isolatePath + "'",
		},
	}

	_, err := c.sysUtils.Execute(isolatePath, "--version")
	if err != nil {
		// TODO log error
	}

	printPretty(err == nil, msg)
	return err == nil
}

// a helper function to print pretty message specific to groundcheck
func printPretty(ok bool, msg *message) {
	if ok {
		color.Green("[ok] %s", msg.success)
		return
	}

	color.HiRed("[  ] %s", msg.fail)
	for _, s := range msg.solve {
		color.Red("     - %s", s)
	}
}
