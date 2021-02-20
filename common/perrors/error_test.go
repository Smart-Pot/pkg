package perrors_test

import (
	"testing"

	"github.com/Smart-Pot/pkg/common/perrors"
	"github.com/stretchr/testify/assert"
)


func TestNew(t *testing.T) {
	var err error = perrors.New("example_message", 105)
	assert.NotNil(t,err)
	assert.NotEqual(t,err.Error(),"")
}


func TestParseError(t *testing.T) {
	const m = "example_message"
	const c = 303
	var err error = perrors.New(m,c)
	pe,ok := err.(*perrors.Error)
	assert.True(t,ok)
	assert.Equal(t,m,pe.Error())
	assert.Equal(t,c,pe.Code())
}