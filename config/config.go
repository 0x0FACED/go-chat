package config

import "time"

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
	MaxTries    int           `toml:"max_tries"`
}

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
	Redis    RedisConfig    `toml:"redis"`
}
