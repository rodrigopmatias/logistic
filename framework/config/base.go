package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	JwtSecret string `json:"jwtSecret"`
	Addr      string `json:"addr"`
	Port      int64  `json:"port"`
	DbDialect string `json:"dbDialect"`
	DbDSN     string `json:"dbDSN"`
}

func (config *Config) Merge(base *Config) {
	if config.JwtSecret == "" && base.JwtSecret != "" {
		config.JwtSecret = base.JwtSecret
	}

	if config.Addr == "" && base.Addr != "" {
		config.Addr = base.Addr
	}

	if config.Port == 0 && base.Port != 0 {
		config.Port = base.Port
	}
}

var defaultConfig *Config
var config *Config

func New() *Config {
	if config == nil {
		defaultConfig = &Config{
			JwtSecret: "secret",
			Addr:      "127.0.0.1",
			Port:      4000,
			DbDialect: "sqlite",
			DbDSN:     "db.sqlite",
		}

		configFile := os.Getenv("APP_CONFIG")
		if configFile == "" {
			configFile = "app.config.json"
		}

		log.Printf("Loading configuration from %s", configFile)
		data, err := ioutil.ReadFile(configFile)

		config = &Config{}
		if err == nil {
			json.Unmarshal(data, &config)
		}

		config.Merge(defaultConfig)
	}

	return config
}
