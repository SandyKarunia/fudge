package utils

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"strings"
)

// Path contains methods to return path for various things
//go:generate mockery -name=Path
type Path interface {
	// BoxDir returns directory for sandbox
	BoxDir(sandboxID uint32) string

	// FudgeDir returns directory for fudge program
	FudgeDir() string
}

type pathImpl struct {
	system System
	sdkOS  sdk.OSFunctions
}

func (p *pathImpl) BoxDir(sandboxID uint32) string {
	return fmt.Sprintf("/var/local/lib/isolate/%d/box/", sandboxID)
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
