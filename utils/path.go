package utils

import (
	"github.com/sandykarunia/fudge/sdk"
	"strings"
)

// Path contains methods to return path for various things
//go:generate mockery -name=Path
type Path interface {
	// Isolate returns path to isolate binary
	IsolateBinary() string

	// BoxDir returns path prefix for the sandbox's box
	BoxDir() string

	// FudgeDir returns directory for fudge program
	FudgeDir() string
}

type pathImpl struct {
	system System
	sdkOS  sdk.OSFunctions
}

func (p *pathImpl) IsolateBinary() string {
	return "/usr/local/bin/isolate"
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

	return homeDir + ".fudge/"
}

// ProvidePath ...
func ProvidePath(system System, sdkOS sdk.OSFunctions) Path {
	return &pathImpl{
		system: system,
		sdkOS:  sdkOS,
	}
}
