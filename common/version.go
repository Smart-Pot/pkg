package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidVersionStr = errors.New("version string is not valid")
)

type version struct {
	major int
	minor int
	patch int
}

type Version interface {
	Major() int
	Minor() int
	Patch() int
	String() string
	Compare(v Version) int
}

func NewVersion(major, minor, patch int) Version {
	return &version{major, minor, patch}
}

func NewVersionFromString(v string) (Version, error) {
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
	return NewVersion(vv[0], vv[1], vv[2]), nil
}

func (v *version) Major() int {
	return v.major
}

func (v *version) Minor() int {
	return v.minor
}

func (v *version) Patch() int {
	return v.patch
}

func (v *version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

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
