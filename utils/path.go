package utils

// Path contains methods to return path for various things
//go:generate mockery -name=Path
type Path interface {
	// Isolate returns path to isolate binary
	IsolateBinary() string
}

type pathImpl struct {
	system System
}

func (p *pathImpl) IsolateBinary() string {
	fudgeDir := p.system.GetFudgeDir()
	return fudgeDir + "isolate"
}

// ProvidePath ...
func ProvidePath(system System) Path {
	return &pathImpl{system: system}
}
