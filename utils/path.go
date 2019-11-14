package utils

import (
	"github.com/sandykarunia/fudge/sdk"
	"strings"
)

// Path contains methods to return path for various things
//go:generate mockery -name=Path
type Path interface {
	// BoxDir returns path prefix for the sandbox's box
	BoxDir() string

	// FudgeDir returns directory for fudge program
	FudgeDir() string
}

type pathImpl struct {
	system System
	sdkOS  sdk.OSFunctions
}

func (p *pathImpl) BoxDir() string {
	return "/var/local/lib/isolate/"
}

func (p *pathImpl) FudgeDir() string {
	homeDir, _ := p.sdkOS.UserHomeDir()

	// if it doesn't end with "/", add it
	if !strings.HasSuffix(homeDir, "/") {
		homeDir += "/"
	}

	fudgeDir := homeDir + ".fudge/"

	// if fudge dir doesn't exist, then create it
	_, _ = p.system.Execute("mkdir", "-p", fudgeDir)

	return fudgeDir
}

// ProvidePath ...
func ProvidePath(system System, sdkOS sdk.OSFunctions) Path {
	return &pathImpl{
		system: system,
		sdkOS:  sdkOS,
	}
}
