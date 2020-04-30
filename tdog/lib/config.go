package lib

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	File string
}

func connect(c *Config) (err error) {
	path, _ := os.Getwd()
	path += "/config"
	file := "app"
	if c != nil {
		file = c.File
	}
	if len(file) < 1 {
		file = "app"
	}

	viper.SetConfigName(file)
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	return
}

func (c *Config) Get(key string) (value interface{}) {
	err := connect(c)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return nil
	}

	if len(key) < 1 {
		value = ""
		return
	}

	value = viper.Get(key)
	if value == nil {
		newConf := new(Config)
		match := strings.Split(key, ".")
		if len(match) > 1 {
			newConf.File = match[0]
			match = match[1:]
		}
		err = connect(newConf)
		if err != nil {
			logger := Logger{Level: 0, Key: "error"}
			logger.New(err.Error())
			return nil
		}
		key = strings.Join(match, ".")
		value = viper.Get(key)
		connect(c)
	}
	return
}
