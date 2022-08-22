package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// VersionRegexp
// vmajor[.minor][.patch]
var VersionRegexp = regexp.MustCompile(`^v(?P<major>[0-9]|[1-9][0-9]|[1-9][0-9][0-9])(?:\.(?P<minor>[0-9]|[1-9][0-9]|[1-9][0-9][0-9]))?(?:\.(?P<patch>[0-9]|[1-9][0-9]|[1-9][0-9][0-9]))?$`)

var ErrSemverVersion = errors.New("version is not semver")

// Version represents the semver of an aggregate.
// v0~999[.0~999][.0~999]
// vmajor[.minor][.patch]
type Version string

// Validate checks if the v is semver
func (v Version) Validate() error {
	if !VersionRegexp.MatchString(string(v)) {
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
func (v Version) IntVersion() IntVersion {
	items := VersionRegexp.FindStringSubmatch(string(v))
	if len(items) != 4 {
		return 0
	}
	var ver int
	n, _ := strconv.Atoi(items[1])
	ver += n * 1e6
	n, _ = strconv.Atoi(items[2])
	ver += n * 1e3
	n, _ = strconv.Atoi(items[3])
	ver += n
	return IntVersion(ver)
}

func (v *Version) Scan(value any) error {
	switch val := value.(type) {
	case int:
		ver, err := ParseIntVersion(val)
		if err != nil {
			return err
		}
		*v = ver
	case string:
		if err := Version(val).Validate(); err != nil {
			return err
		}
		*v = Version(val)
	default:
		return fmt.Errorf("unexpect Version type: %s", reflect.TypeOf(value))
	}
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
		return "v0", fmt.Errorf("max version is 999999999")
	}
	if version < 0 {
		return "v0", fmt.Errorf("min version is 0")
	}
	buf := strings.Builder{}
	patch := version % 1e3
	version /= 1e3
	minor := version % 1e3
	version /= 1e3
	major := version % 1e3
	buf.WriteByte('v')
	buf.WriteString(strconv.Itoa(major))
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(minor))
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(patch))
	return Version(buf.String()), nil
}

// IntVersion Int version
// 0~1000000000
type IntVersion int

func (v IntVersion) Version() Version {
	ver, _ := ParseIntVersion(int(v))
	return ver
}

// Validate checks if the v is semver
func (v IntVersion) Validate() error {
	if v > 999999999 {
		return fmt.Errorf("max version is 999999999")
	}
	if v < 0 {
		return fmt.Errorf("min version is 0")
	}
	return nil
}

func (v IntVersion) String() string {
	ver, _ := ParseIntVersion(int(v))
	return string(ver)
}

func (v *IntVersion) Scan(value any) error {
	switch val := value.(type) {
	case int:
		*v = IntVersion(val)
		if err := v.Validate(); err != nil {
			return err
		}
	case string:
		ver := Version(val)
		if err := ver.Validate(); err != nil {
			return err
		}
		*v = ver.IntVersion()
	default:
		return fmt.Errorf("unexpect Version type: %s", reflect.TypeOf(value))
	}
	return nil
}

func (v IntVersion) Value() (driver.Value, error) {
	if err := v.Validate(); err != nil {
		return nil, err
	}
	return int(v), nil
}
