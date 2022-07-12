package types

import (
	"errors"
	"regexp"
)

var versionRegexp = regexp.MustCompile(`^v\d+(\.\d+){0,2}$`)

// Version represents the semver of an aggregate
type Version string

// Validate checks if the v is semver
func (v Version) Validate() error {
	if !versionRegexp.MatchString(string(v)) {
		return errors.New("version is not semver")
	}
	return nil
}

// Int the v is integer semver
func (v Version) Int() int {
	return 0
}
