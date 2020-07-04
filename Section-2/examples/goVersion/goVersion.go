package golang

import "runtime"

// Version returns go version
func Version() string {
	return runtime.Version()
}
