package conf

import (
	"errors"

	"github.com/robfig/config"
)

var configFile string
var c *config.Config

func SetFile(f string) error {
	configFile = f
	_, err := GetConf()
	if err != nil {
		return err
	}
	return nil
}
func GetConf() (*config.Config, error) {
	if c != nil {
		return c, nil
	}
	tmpc, err := config.ReadDefault(configFile)
	if err != nil {
		return nil, err
	}
	c = tmpc
	return c, nil
}

func GetString(section string, option string) (value string, err error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	return c.String(section, option)
}
func GetInt(section string, option string) (value int, err error) {
	if c == nil {
		return 0, errors.New("config is nil")
	}
	return c.Int(section, option)
}
func GetInt32(section string, option string) (value int32, err error) {
	i, err := GetInt(section, option)
	return int32(i), err
}
func GetInt64(section string, option string) (value int64, err error) {
	i, err := GetInt(section, option)
	return int64(i), err
}

func GetFloat(section string, option string) (value float64, err error) {
	if c == nil {
		return 0, errors.New("config is nil")
	}
	return c.Float(section, option)
}
func GetBool(section string, option string) (value bool, err error) {
	if c == nil {
		return false, errors.New("config is nil")
	}
	return c.Bool(section, option)
}
