/*
Package pkg implements configuration utilities by reading from a config file.
config file path and file extention can be configured by user. use `pkg.ConfigOptions`


Usage:

	package main

	import (
		...
		"github.com/Smart-Pot/pkg/"
		...
	)

	func main(){
		...

		if err := pkg.Config.ReadConfig(); err != nil {
			// Handle error
		}
		...

		address := pkg.Config.Server.Address
	}
*/
package pkg

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type databaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Net      string `json:"net"`
	Addr     string `json:"addr"`
	DBName   string `json:"dbName"`
}

type serverConfig struct {
	Address string `json:"address"`
}
type config struct {
	Database    databaseConfig `json:"database"`
	Server      serverConfig   `json:"server"`
	AMQPAddress string         `json:"amqpAdress"`
}

type configOptions struct {
	// BaseDir is base directory that config file is located
	BaseDir string
	// ConfigType is file type of configuration file.
	// it could be: json, yaml, yml, etc.
	ConfigType string
}

// ConfigOptions is options for reading config file process
var ConfigOptions configOptions = initConfigOptions()

func initConfigOptions() configOptions {
	wd, _ := os.Getwd()
	return configOptions{
		BaseDir:    filepath.Join(wd, "config"),
		ConfigType: "yml",
	}
}

// Config represent  the config file in runtime
var Config config

// fillDefaults set default values for configuration
func (c *config) fillDefaults() {
	c.Server = serverConfig{":3000"}
	c.Database = databaseConfig{
		User:     "",
		Password: "",
		Net:      "",
		Addr:     "mongodb://localhost:27017",
		DBName:   "test",
	}
	c.AMQPAddress = "amqp://guest:guest@localhost:5672"
}

// ReadConfig reads configurations from given file with given options in ConfigOptions
func (c *config) ReadConfig() error {

	c.fillDefaults()
	C := &c

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType(ConfigOptions.ConfigType)
	v.AddConfigPath(ConfigOptions.BaseDir)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&C); err != nil {
		return err
	}
	return nil
}
