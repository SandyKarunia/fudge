package utils

import "os"

// IsSudo returns true if this program run with sudo
func IsSudo() bool {
	if os.Geteuid() != 0 || len(os.Getenv("SUDO_UID")) == 0 ||
		len(os.Getenv("SUDO_GID")) == 0 || len(os.Getenv("SUDO_USER")) == 0 {
		return false
	}

	return true
}
