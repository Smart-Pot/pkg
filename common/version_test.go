package common

import (
	"testing"
)

func TestVersion(t *testing.T) {

	t.Run("Create Version", func(t *testing.T) {
		v := NewVersion(1, 2, 15)
		if v.String() != "1.2.15" {
			t.Error("want", "1.2.15", " got", v.String())
			t.FailNow()
		}
	})

	t.Run("New Version From String", func(t *testing.T) {
		str := "1.5.3"
		sv := NewVersion(1, 5, 3)
		sver, err := NewVersionFromString(str)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if str != sv.String() {
			t.Error("want", str, "got", sv.String())
			t.FailNow()
		}
		if sver.Compare(sv) == 0 {
			t.Error("Not equal", sver, sv)
			t.FailNow()
		}

	})

	t.Run("Compare Version", func(t *testing.T) {
		big := NewVersion(1, 5, 3)
		small := NewVersion(1, 4, 3)
		if big.Compare(small) != 1 {
			t.FailNow()
		}
		if small.Compare(big) != -1 {
			t.FailNow()
		}
		if big.Compare(big) != 0 {
			t.FailNow()
		}

	})

}
