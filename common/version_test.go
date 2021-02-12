package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestVersion_String(t *testing.T){
	v := NewVersion(1,2,15)
	assert.Equal(t,"1.2.15",v.String())
}

func TestVersion_Compare(t *testing.T){
	
	big := NewVersion(1,2,15)
	small := NewVersion(1,0,1)
	
	assert.Equal(t,1,big.Compare(small))
	assert.Equal(t,0,big.Compare(big))
	assert.Equal(t,-1,small.Compare(big))
}


func TestNewVersionFromString(t *testing.T) {
	versionStr := "1.2.15"
	v,err := NewVersionFromString(versionStr)
	assert.Nil(t,err)
	assert.Equal(t,v.Major(),1)
	assert.Equal(t,v.Minor(),2)
	assert.Equal(t,v.Patch(),15)
}

