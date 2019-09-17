package sandbox

import "errors"

// ErrMoeIsNotReady means the sandbox is not ready
var ErrMoeIsNotReady = errors.New("moe sandbox is not ready yet, please call Initialize() before use")

// ErrMoeIsGone means the sandbox is gone / destroyed
var ErrMoeIsGone = errors.New("moe sandbox is gone / destroyed and should not be used anymore")

const (
	// MoeDefaultConfigSource is the source location of the default configuration
	MoeDefaultConfigSource = "./bin/moe/default.cf"
	// MoeDefaultConfigDestination is the default location of isolate configuration
	MoeDefaultConfigDestination = "/usr/local/etc/isolate"
)
