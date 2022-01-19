package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	fileName          string
	ConsumerKey       string
	ConsumerKeySecret string
	AccessToken       string
	AccessTokenSecret string
}

type ConfigJson struct {
	ConsumerKey       string `json:'ConsumerKey`
	ConsumerKeySecret string `json:'ConsumerKeySecret`
	AccessToken       string `json:'AccessToken`
	AccessTokenSecret string `json:'AccessTokenSecret`
}

func NewConfig() *Config {
	return &Config{fileName: "config.json"}
}

func (c *Config) Save(cfg ConfigJson) error {
	f, err := os.Create(c.fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Load() error {
	f, err := os.Open(c.fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	var cfg ConfigJson
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return err
	}

	c.ConsumerKey = cfg.ConsumerKey
	c.ConsumerKeySecret = cfg.ConsumerKeySecret
	c.AccessToken = cfg.AccessToken
	c.AccessTokenSecret = cfg.AccessTokenSecret
	return nil
}
