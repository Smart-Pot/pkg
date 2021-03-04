package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestVersion_String(t *testing.T){
	v := New(1,2,15)
	assert.Equal(t,"1.2.15",v.String())
}

func TestVersion_Compare(t *testing.T){
	
	big := New(1,2,15)
	small := New(1,0,1)
	
	assert.Equal(t,1,big.Compare(small))
	assert.Equal(t,0,big.Compare(big))
	assert.Equal(t,-1,small.Compare(big))
}



func TestNewVersionFromString(t *testing.T) {
	versionStr := "1.2.15"
	v,err := FromString(versionStr)
	assert.Nil(t,err)
	assert.Equal(t,v.Major(),1)
	assert.Equal(t,v.Minor(),2)
	assert.Equal(t,v.Patch(),15)
}

func TestReadFromFile(t *testing.T) {
/*

	type VC struct {
		Version Version `json:"version"`
	}
	t.Run("From JSON",func(t *testing.T){
		f := `{"version": "1.0.1" }`
		var vc VC
		// var vv Version
		assert.Nil(t,json.Unmarshal([]byte(f),&vc))
		assert.NotNil(t,vc.Version)

		v, err := FromString("1.0.1")
		assert.Nil(t,err)
		t.Log(v,vc)
	})

*/
}