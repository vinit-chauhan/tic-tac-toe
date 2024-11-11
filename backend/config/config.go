package config

import (
	"os"

	"github.com/vinit-chauhan/tic-tac-toe/utils/logger"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DbConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Server   ServerConfig `yaml:"server"`
	Database DbConfig     `yaml:"database"`
	Redis    RedisConfig  `yaml:"redis"`
}

func Default() Config {
	return Config{
		Server:   ServerConfig{Port: 8080, Host: "localhost"},
		Database: DbConfig{},
		Redis:    RedisConfig{},
	}
}

func Load(path string) (Config, error) {
	var conf Config

	buff, err := os.ReadFile(path)
	if err != nil {
		logger.Error("unable to read config file", "Load", err)
		return Config{}, err
	}

	err = yaml.Unmarshal(buff, &conf)
	if err != nil {
		logger.Error("unable to marshal the yaml file", "Load", err)
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
