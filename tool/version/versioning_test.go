package versioning

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMatch(t *testing.T) {
	r, err := http.NewRequest("GET", "google.com", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	r.Header.Set(apiVersionKey, "1.5.3")
	testCases := []struct {
		query string
		res   bool
	}{
		{
			query: "> 1.0.0",
			res:   true,
		},
		{
			query: "= 1.5.3",
			res:   true,
		},
		{
			query: "< 1.0.0",
			res:   false,
		},
		{
			query: ">= 1.5.3",
			res:   true,
		},
	}

	for i, testCase := range testCases {
		res, err := Match(r, testCase.query)
		if err != nil {
			t.Error(fmt.Errorf("Failed on test case #%d : %s", i, err))
			t.FailNow()
		}
		if res != testCase.res {
			t.Error("Want", testCase.res, "got", res)
			t.Fail()
		}
	}

}
