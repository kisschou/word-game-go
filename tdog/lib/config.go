package lib

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		File     string
		Callback *Callback
	}

	Callback struct {
		key                  string
		isExists             bool
		rawData              interface{}
		stringData           string
		intData              int
		boolData             bool
		intSlice             []int
		stringMap            map[string]interface{}
		stringMapString      map[string]string
		stringMapStringSlice map[string][]string
		stringSlice          []string
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

	if !viper.InConfig(key) {
		newConf := new(Config)
		match := strings.Split(key, ".")
		if len(match) > 1 {
			newConf.File = match[0]
			file = match[0]
			match = match[1:]
		}
		beginConn(newConf)
		key = strings.Join(match, ".")
		if !viper.InConfig(key) {
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
	beginConn(c)
	callback := c.Callback.New(key)
	if file == "" || key == "" {
		callback.isExists = false
	} else {
		callback.isExists = true
		callback.Fill()
	}
	c.Callback = callback
	endConn()
	return c
}

func (c *Config) IsExists() bool {
	return c.Callback.isExists
}

func (c *Config) RawData() interface{} {
	return c.Callback.rawData
}

func (c *Config) String() string {
	return c.Callback.stringData
}

func (c *Config) Int() int {
	return c.Callback.intData
}

func (c *Config) Bool() bool {
	return c.Callback.boolData
}

func (c *Config) IntSlice() []int {
	return c.Callback.intSlice
}

func (c *Config) StringMap() map[string]interface{} {
	return c.Callback.stringMap
}

func (c *Config) StringMapString() map[string]string {
	return c.Callback.stringMapString
}

func (c *Config) StringMapStringSlice() map[string][]string {
	return c.Callback.stringMapStringSlice
}

func (c *Config) StringSlice() []string {
	return c.Callback.stringSlice
}

func (c *Callback) New(key string) *Callback {
	return &Callback{
		key:                  key,
		rawData:              nil,
		stringData:           "",
		intData:              0,
		boolData:             false,
		intSlice:             nil,
		stringMap:            nil,
		stringMapString:      nil,
		stringMapStringSlice: nil,
		stringSlice:          nil,
	}
}

func (c *Callback) Fill() {
	c.rawData = viper.Get(c.key)
	c.stringData = viper.GetString(c.key)
	c.intData = viper.GetInt(c.key)
	c.boolData = viper.GetBool(c.key)
	c.intSlice = viper.GetIntSlice(c.key)
	c.stringMap = viper.GetStringMap(c.key)
	c.stringMapString = viper.GetStringMapString(c.key)
	c.stringMapStringSlice = viper.GetStringMapStringSlice(c.key)
	c.stringSlice = viper.GetStringSlice(c.key)
}
