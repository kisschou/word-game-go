package lib

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		File string
		Key  string
	}
)

func beginConn(c *Config) {
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
	err := viper.ReadInConfig()
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
	}
}

func action(sourceKey string) (file string, key string) {
	conf := new(Config)
	file = ""
	key = ""

	if len(sourceKey) < 1 {
		return
	}

	beginConn(conf)
	file = conf.File
	key = sourceKey

	if !viper.IsSet(key) {
		newConf := new(Config)
		match := strings.Split(key, ".")
		if len(match) > 1 {
			newConf.File = match[0]
			file = match[0]
			match = match[1:]
		}
		beginConn(newConf)
		key = strings.Join(match, ".")
		if !viper.IsSet(key) {
			file = ""
			key = ""
		}
	}

	endConn()
	return
}

func endConn() {
	var c *Config
	beginConn(c)
}

func (c *Config) Get(sourceKey string) *Config {
	file, key := action(sourceKey)
	c.File = file
	c.Key = key
	endConn()
	return c
}

func (c *Config) IsExists() bool {
	isExists := false
	beginConn(c)
	if c.File == "" || c.Key == "" {
		isExists = false
	} else {
		isExists = true
	}
	endConn()
	return isExists
}

func (c *Config) RawData() interface{} {
	return viper.Get(c.Key)
}

func (c *Config) String() string {
	return viper.GetString(c.Key)
}

func (c *Config) Int() int {
	return viper.GetInt(c.Key)
}

func (c *Config) Bool() bool {
	return viper.GetBool(c.Key)
}

func (c *Config) IntSlice() []int {
	return viper.GetIntSlice(c.Key)
}

func (c *Config) StringMap() map[string]interface{} {
	return viper.GetStringMap(c.Key)
}

func (c *Config) StringMapString() map[string]string {
	return viper.GetStringMapString(c.Key)
}

func (c *Config) StringMapStringSlice() map[string][]string {
	return viper.GetStringMapStringSlice(c.Key)
}

func (c *Config) StringSlice() []string {
	return viper.GetStringSlice(c.Key)
}
