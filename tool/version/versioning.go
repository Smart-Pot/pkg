package versioning

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Smart-Pot/pkg/common/version"
)

const (
	apiVersionKey = "X-API-VERSION"
)

var (
	ErrInvalidQuery      = errors.New("invalid match query")
	ErrVersionNotFound   = errors.New("api version not found in request")
	ErrInvalidAPIVersion = errors.New("api version is invalid in request")
)

var acceptMap = map[string][]int{
	"<":  {-1},
	"<=": {-1, 0},
	">":  {1},
	">=": {1, 0},
	"=":  {0},
}

func Match(r *http.Request, query string) (bool, error) {
	avstr := r.Header.Get(apiVersionKey)

	if avstr == "" {
		return false, ErrVersionNotFound
	}
	apiv, err := version.FromString(avstr)
	if err != nil {
		return false, ErrInvalidAPIVersion
	}

	accepts, qver, err := parseQuery(query)
	if err != nil {
		return false, err
	}
	c := apiv.Compare(qver)
	for _, accept := range accepts {
		if accept == c {
			return true, nil
		}
	}
	return false, nil
}

func parseQuery(q string) ([]int, version.Version, error) {
	parts := strings.Split(q, " ")
	if len(parts) != 2 {
		return nil, nil, ErrInvalidQuery
	}

	ac, ok := acceptMap[parts[0]]
	if !ok {
		return nil, nil, ErrInvalidQuery
	}
	qv, err :=  version.FromString(parts[1])
	if err != nil {
		return nil, nil, ErrInvalidQuery
	}

	return ac, qv, nil
}
