package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const RunEnv = "dev"

func NewConfig(name interface{}) *config {
	value, ok := name.(string)
	if !ok {
		value = RunEnv
	}
	c := new(config)
	c.viper = viper.New()
	c.initConfig(value)
	return c
}

type config struct {
	viper *viper.Viper
}

func (c *config) initConfig(name string) {
	c.viper.SetConfigName(name)                    // name of config file (without extension)
	c.viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	c.viper.AddConfigPath("./app/config")          // path to look for the config file in
	if err := c.viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (c *config) GetMapString(name string) map[string]string {
	return c.viper.GetStringMapString(name)
}

func (c *config) GetValue(name string) string {
	return c.viper.GetString(name)
}
