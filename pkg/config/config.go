package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConsumerKey       string
	ConsumerKeySecret string
	AccessToken       string
	AccessTokenSecret string
	fileName          string
	filePath          string
	fileType          string
}

func NewConfig() *Config {
	return &Config{
		fileName: "config.json",
		filePath: "./",
		fileType: "json",
	}
}

func (c *Config) Save() error {
	viper.Set("ConsumerKey", c.ConsumerKey)
	viper.Set("ConsumerKeySecret", c.ConsumerKeySecret)
	viper.Set("AccessToken", c.AccessToken)
	viper.Set("AccessTokenSecret", c.AccessTokenSecret)

	viper.SetConfigType(c.fileType)
	return viper.WriteConfigAs(c.filePath + c.fileName)
}

func (c *Config) Load() (*Config, error) {
	viper.SetConfigName(c.fileName)
	viper.SetConfigType(c.fileType)
	viper.AddConfigPath(c.filePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
