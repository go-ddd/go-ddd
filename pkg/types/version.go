package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

// versionRegexp
// vmajor[.minor][.patch]
var versionRegexp = regexp.MustCompile(`^v(?P<major>[0-9]|[1-9][0-9]|[1-9][0-9][0-9])(?:\.(?P<minor>[0-9]|[1-9][0-9]|[1-9][0-9][0-9]))?(?:\.(?P<patch>[0-9]|[1-9][0-9]|[1-9][0-9][0-9]))?$`)

var ErrSemverVersion = errors.New("version is not semver")

// Version represents the semver of an aggregate.
// v0~999[.0~999][.0~999]
// vmajor[.minor][.patch]
type Version string

// Validate checks if the v is semver
func (v Version) Validate() error {
	if !versionRegexp.MatchString(string(v)) {
		return ErrSemverVersion
	}
	return nil
}

func (v Version) String() string {
	return string(v)
}

// IntVersion the v is integer semver
// v0~999[.0~999][.0~999]
// vmajor[.minor][.patch]
// int version is vmajor*1e6+minor*1e3+patch
func (v Version) IntVersion() int {
	items := versionRegexp.FindStringSubmatch(string(v))
	if len(items) != 4 {
		return -1
	}
	var ver int
	n, _ := strconv.Atoi(items[1])
	ver += n * 1e6
	n, _ = strconv.Atoi(items[2])
	ver += n * 1e3
	n, _ = strconv.Atoi(items[3])
	ver += n
	return ver
}

func (v *Version) Scan(value any) error {
	num, ok := value.(int)
	if !ok {
		return fmt.Errorf("want to int type, but got type: %s", reflect.TypeOf(value))
	}
	ver, err := ParseIntVersion(num)
	if err != nil {
		return err
	}
	*v = ver
	return nil
}

func (v Version) Value() (driver.Value, error) {
	if err := v.Validate(); err != nil {
		return nil, err
	}
	return v.IntVersion(), nil
}

func ParseIntVersion(version int) (Version, error) {
	if version > 999999999 {
		return "", fmt.Errorf("max version is 999999999")
	}
	if version < 0 {
		return "", fmt.Errorf("min version is 0")
	}
	return "", nil
}
