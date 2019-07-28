package version

import (
	"fmt"
)

// The following variables are private fields and should be set during compilation. See Makefile
var (
	// The current version of the binary
	binaryVersion = "0.0.0-unset"
	imageName     = "jenkinsxio/jx-app-jacoco"
)

// GetVersion returns the version of this binary.
func GetVersion() string {
	return binaryVersion
}

// GetFQImage returns the fully qualified image name to be used within the pipeline.
func GetFQImage() string {
	return fmt.Sprintf("%s:%s", imageName, GetVersion())
}
