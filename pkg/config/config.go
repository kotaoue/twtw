package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	fileName    string
	bearerToken string
}

type configJson struct {
	BearerToken string `json:'BearerToken`
}

func NewConfig() *Config {
	return &Config{fileName: "config.json"}
}

func (c *Config) Save(token string) error {
	f, err := os.Create(c.fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.MarshalIndent(configJson{BearerToken: token}, "", "    ")
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

	var cfg configJson
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return err
	}

	c.bearerToken = cfg.BearerToken
	return nil
}
