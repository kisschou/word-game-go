package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Path string
	Name string
	Type string
}

func connect(conf *Config) (err error) {
	var configPath, configName, configType string

	if len(conf.Path) > 0 {
		configPath = conf.Path
	} else {
		configPath, _ = os.Getwd()
		configPath += "/config"
	}

	if len(conf.Name) > 0 {
		configName = conf.Name
	} else {
		configName = "config"
	}

	if len(conf.Type) > 0 {
		configType = conf.Type
	} else {
		configType = "toml"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err = viper.ReadInConfig()
	return
}

func (conf *Config) Get(key string) (value interface{}) {
	err := connect(conf)
	if err != nil {
		return nil
	}

	value = viper.Get(key)
	if value == nil {
		newConf := new(Config)
		match := strings.Split(key, ".")
		if len(match) > 1 {
			newConf.Name = match[0]
			match = match[1:]
		}
		err = connect(newConf)
		if err != nil {
			return nil
		}
		key = strings.Join(match, ".")
		value = viper.Get(key)
		connect(conf)
	}

	return
}
