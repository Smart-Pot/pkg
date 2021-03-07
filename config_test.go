package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConfig = config{
	Server: serverConfig{
		Address: ":3000",
	},
	Database: databaseConfig{
		Addr:   "example.url",
		DBName: "exampledb",
	},
	AMQPAddress: "amqp.address",
}

func TestMain(m *testing.M) {
	// Set basedir
	wd, _ := os.Getwd()
	ConfigOptions.BaseDir = filepath.Join(wd, "test_files", "config")
	c := m.Run()
	os.Exit(c)
}

func TestReadConfigJSON(t *testing.T) {
	// Set config options for json
	ConfigOptions.ConfigType = "json"
	if err := Config.ReadConfig(); err != nil {
		assert.Nil(t, err)
	}

	// Check fields are correct
	assert.Equal(t, testConfig.Server.Address, Config.Server.Address)
	assert.Equal(t, testConfig.Database.Addr, Config.Database.Addr)
	assert.Equal(t, testConfig.Database.DBName, Config.Database.DBName)
	assert.Equal(t, testConfig.AMQPAddress, Config.AMQPAddress)
}


func TestReadConfigYAML(t *testing.T) {
	// Set config options for json
	ConfigOptions.ConfigType = "json"
	if err := Config.ReadConfig(); err != nil {
		assert.Nil(t, err)
	}

	// Check fields are correct
	assert.Equal(t, testConfig.Server.Address, Config.Server.Address)
	assert.Equal(t, testConfig.Database.Addr, Config.Database.Addr)
	assert.Equal(t, testConfig.Database.DBName, Config.Database.DBName)
	assert.Equal(t, testConfig.AMQPAddress, Config.AMQPAddress)
}


func TestReadConfigToStruct(t *testing.T) {
	var v struct{
		Test string
		Hello string
	}
	err := readConfigToStruct(ConfigOptions.BaseDir,"custom","json",&v)
	assert.Nil(t,err)
	assert.Equal(t,"Message",v.Test)
	assert.Equal(t,"World",v.Hello)
}

func TestUnmarshalConfigFromFile(t *testing.T) {
	var v struct{
		Test string
		Hello string
	}
	err := Config.UnmarshalConfigFromFile(ConfigOptions.FromBaseDir("custom.json"),&v)
	assert.Nil(t,err)
	assert.Equal(t,"Message",v.Test)
	assert.Equal(t,"World",v.Hello)
}

func TestSplitFilename(t *testing.T) {
	tests := []struct {
		dir string
		filename string
		ext string
	} {
		{"the/base/dir/","filename",".json"},
	}

	for _,ti := range tests {
		pp := ti.dir + ti.filename + ti.ext
		d,f,e := splitFilename(pp)
		assert.Equal(t,ti.dir,d)
		assert.Equal(t,ti.filename,f)
		assert.Equal(t,ti.ext[1:],e)
	}
}