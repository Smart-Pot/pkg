package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)


var testConfig =  config{
	Server: serverConfig{
		Address: ":3000",
	},
	Database:databaseConfig{
		Addr: "example.url"	,
		DBName: "exampledb",
	},
	AMQPAddress: "amqp.address",
}

func TestMain(m *testing.M) {
	// Set basedir
	wd,_ := os.Getwd()
	ConfigOptions.BaseDir = filepath.Join(wd,"test_files","config")
	c := m.Run()
	os.Exit(c)
}

func TestReadConfigJSON(t *testing.T){
	// Set config options for json
	ConfigOptions.ConfigType = "json"
	if err := Config.ReadConfig(); err != nil {
		assert.Nil(t,err)
	}

	// Check fields are correct
	assert.Equal(t,testConfig.Server.Address,Config.Server.Address)
	assert.Equal(t,testConfig.Database.Addr,Config.Database.Addr)
	assert.Equal(t,testConfig.Database.DBName,Config.Database.DBName)
	assert.Equal(t,testConfig.AMQPAddress,Config.AMQPAddress)
}


func TestReadConfigYAML(t *testing.T){
	// Set config options for json
	ConfigOptions.ConfigType = "json"
	if err := Config.ReadConfig(); err != nil {
		assert.Nil(t,err)
	}

	// Check fields are correct
	assert.Equal(t,testConfig.Server.Address,Config.Server.Address)
	assert.Equal(t,testConfig.Database.Addr,Config.Database.Addr)
	assert.Equal(t,testConfig.Database.DBName,Config.Database.DBName)
	assert.Equal(t,testConfig.AMQPAddress,Config.AMQPAddress)
}