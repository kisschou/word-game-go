package lib

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Path string
	File string
	Key  string
}

func connect(c *Config) (err error) {
	path, _ := os.Getwd()
	path += "/config"
	file := "app"

	if len(c.Path) > 0 {
		path = c.Path
	}

	if len(c.File) > 0 {
		file = c.File
	}

	viper.SetConfigName(file)
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	return
}

func (c *Config) Get(keys ...string) (value interface{}) {
	err := connect(c)
	key := ""
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return nil
	}

	if (len(c.Key) < 1) && len(keys) != 1 {
		value = ""
		return
	}

	if len(c.Key) > 0 {
		key = c.Key
	} else {
		key = keys[0]
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
