package config

import (
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
	"time"
)

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type DatabaseConfig struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	Username   string `toml:"username"`
	Password   string `toml:"password"`
	DBName     string `toml:"dbname"`
	DriverName string `toml:"driver"`
}

type RedisConfig struct {
	Host        string        `toml:"host"`
	Port        int           `toml:"port"`
	Network     string        `toml:"network"`
	Username    string        `toml:"username"`
	Password    string        `toml:"password"`
	DialTimeout time.Duration `toml:"dial_timeout"`
	MaxRetries  int           `toml:"max_tries"`
}

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
	Redis    RedisConfig    `toml:"redis"`
}

func Load() (*Config, error) {
	cfgFile, err := os.Open("./config/config.toml")
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()
	cfgBytes, err := io.ReadAll(cfgFile)
	if err != nil {
		return nil, err
	}
	var config *Config
	if err := toml.Unmarshal(cfgBytes, &config); err != nil {
		return nil, err
	}
	return config, nil
}
