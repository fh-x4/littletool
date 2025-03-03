package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	HttpServe  string `json:"http_serve"`
	Log        string `json:"log"`
	MaxBodyLen int    `json:"max_body_len"`
}

var config *Config

func Init(location string) error {
	data, err := os.ReadFile(location)
	if err != nil {
		return err
	}
	config = new(Config)
	if err := json.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}

func Get() *Config {
	return config
}
