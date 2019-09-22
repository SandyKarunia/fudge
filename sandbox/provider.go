package sandbox

import "gitlab.com/sandykarunia/fudge/utils"

// Provider ...
func Provider(path utils.Path) Factory {
	return &factoryImpl{
		path: path,
	}
}
