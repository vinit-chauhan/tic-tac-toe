package config

import (
	"os"

	"github.com/vinit-chauhan/tic-tac-toe/pkg/logger"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DbConfig struct{}

type Config struct {
	Server   ServerConfig `yaml:"server"`
	Database DbConfig     `yaml:"database"`
}

func Default() Config {
	return Config{
		Server:   ServerConfig{Port: 8080, Host: "localhost"},
		Database: DbConfig{},
	}
}

func Load(path string) (Config, error) {
	var conf Config

	buff, err := os.ReadFile(path)
	if err != nil {
		logger.Error("[Load] unable to read config file", err)
		return Config{}, err
	}

	err = yaml.Unmarshal(buff, &conf)
	if err != nil {
		logger.Error("[Load] unable to marshal the yaml file", err)
		return Config{}, err
	}

	return conf, nil
}

func (c *Config) WithServerConfig(s ServerConfig) *Config {
	c.Server = s
	return c
}

func (c *Config) WithDbConfig(d DbConfig) *Config {
	c.Database = d
	return c
}
