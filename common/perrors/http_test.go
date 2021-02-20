package perrors_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Smart-Pot/pkg/common/perrors"
	"github.com/stretchr/testify/assert"
)



var _ http.ResponseWriter = (*mockRW)(nil)

type mockRW struct {
	code int
	resp []byte
	h http.Header
}
func newMockRW() *mockRW {
	return &mockRW{
		resp: []byte{},
		h : map[string][]string{},
	}
}

func (m *mockRW) Header() http.Header {
	return m.h
}

func (m *mockRW) WriteHeader(code int) {
	m.code = code
}

func (m *mockRW) Write(b []byte) (int,error) {
	m.resp = append(m.resp, b...)
	return 0,nil
}

func TestEncodeHTTPError(t *testing.T) {
	const m = "test error bad request"
	const c = 401
	rw := newMockRW()
	ctx := context.TODO()
	perr := perrors.New(m,c)
	perrors.EncodeHTTPError(ctx,perr,rw)

	rj := fmt.Sprintf("{\"code\":%d,\"error\":\"%s\"}\n",c,m)

	assert.Equal(t,c,rw.code)
	assert.Equal(t,rj,string(rw.resp))

}