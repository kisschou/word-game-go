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

func (c *Config) RawData() (data interface{}) {
	beginConn(c)
	data = viper.Get(c.Key)
	endConn()
	return
}

func (c *Config) String() (data string) {
	beginConn(c)
	data = viper.GetString(c.Key)
	endConn()
	return
}

func (c *Config) Int() (data int) {
	beginConn(c)
	data = viper.GetInt(c.Key)
	endConn()
	return
}

func (c *Config) Bool() (data bool) {
	beginConn(c)
	data = viper.GetBool(c.Key)
	endConn()
	return
}

func (c *Config) IntSlice() (data []int) {
	beginConn(c)
	data = viper.GetIntSlice(c.Key)
	endConn()
	return
}

func (c *Config) StringMap() (data map[string]interface{}) {
	beginConn(c)
	data = viper.GetStringMap(c.Key)
	endConn()
	return
}

func (c *Config) StringMapString() (data map[string]string) {
	beginConn(c)
	data = viper.GetStringMapString(c.Key)
	endConn()
	return
}

func (c *Config) StringMapStringSlice() (data map[string][]string) {
	beginConn(c)
	data = viper.GetStringMapStringSlice(c.Key)
	endConn()
	return
}

func (c *Config) StringSlice() (data []string) {
	beginConn(c)
	data = viper.GetStringSlice(c.Key)
	endConn()
	return
}
