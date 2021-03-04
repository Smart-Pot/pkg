// Package version implements bundle of version utilities for versioning APIs
package version

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrInvalidVersionStr :
	ErrInvalidVersionStr = errors.New("version string is not valid")
)




type version struct {
	major int
	minor int
	patch int
}
// Version represents a semantic version
type Version interface {
	Major() int
	Minor() int
	Patch() int
	String() string
	Compare(v Version) int
	json.Unmarshaler
	json.Marshaler
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// New creates a new version
func New(major, minor, patch int) Version {
	return &version{major, minor, patch}
}

// FromString creates a new string from plain version string
// such as: '1.2.5','3.5.4'
func FromString(v string) (Version, error) {
	return newVersionFromString(v)
}
func newVersionFromString(v string) (*version,error) {
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidVersionStr
	}
	vv := make([]int, 3)
	for i, m := range parts {
		mm, err := strconv.ParseInt(m, 10, 32)
		if err != nil {
			return nil, err
		}
		vv[i] = int(mm)
	}
	return &version{vv[0], vv[1], vv[2]}, nil

}
// Major returns a version for incompatible API changes,
func (v *version) Major() int {
	return v.major
}
// Minor returns a version for functionalities in a backwards compatible manner
func (v *version) Minor() int {
	return v.minor
}
// Patch returns a version for backwards compatible bug fixes.
func (v *version) Patch() int {
	return v.patch
}

func (v *version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

// Compare compares two versions and returns a integer for each situtation
// if other > this   returns -1
// if this  > other  returns  1
// if this  = other  returns  0
func (v *version) Compare(ot Version) int {
	r := v.compareNum(v.Major(), ot.Major())
	if r != 0 {
		return r
	}
	r = v.compareNum(v.Minor(), ot.Minor())
	if r != 0 {
		return r
	}
	r = v.compareNum(v.Patch(), ot.Patch())
	if r != 0 {
		return r
	}
	return 0
}

func (v *version) compareNum(x, y int) int {
	if x > y {
		return 1
	}
	if y > x {
		return -1
	}
	return 0
}


func (v *version) UnmarshalJSON(b []byte) error {
	fmt.Println("HEL")
	return v.UnmarshalText(b)
}

func (v *version) MarshalJSON() ([]byte,error) {
	return v.MarshalText()
}

func (v *version) UnmarshalText(b []byte) error {
	s := string(b)
	s = strings.Trim(s,"\"")
	vv,err := newVersionFromString(s)
	if err != nil {
		return err
	}
	v = vv
	return nil
}

func (v *version) MarshalText() ([]byte,error) {
	return []byte(v.String()),nil
}