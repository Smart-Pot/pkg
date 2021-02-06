package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Net      string
		Addr     string
		DBName   string
	}
	Server struct {
		Address string
	}
	AMQPAddress string
}

var ConfigOptions struct {
	BaseDir    string
	ConfigType string
}

func init() {
	wd, _ := os.Getwd()
	ConfigOptions = struct {
		BaseDir    string
		ConfigType string
	}{
		BaseDir:    filepath.Join(wd, "config"),
		ConfigType: "yml",
	}
}

var Config config

func (c *config) fillDefaults() {
	c.Server = struct{ Address string }{":3000"}
	c.Database = struct {
		User     string
		Password string
		Net      string
		Addr     string
		DBName   string
	}{
		User:     "",
		Password: "",
		Net:      "",
		Addr:     "mongodb://localhost:27017",
		DBName:   "test",
	}
	c.AMQPAddress = "amqp://guest:guest@localhost:5672"
}

func (c *config) ReadConfig() error {
	c.fillDefaults()
	C := &c
	viper.SetConfigName("config")
	viper.SetConfigType(ConfigOptions.ConfigType)
	viper.AddConfigPath(ConfigOptions.BaseDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&C); err != nil {
		return err
	}
	return nil
}
