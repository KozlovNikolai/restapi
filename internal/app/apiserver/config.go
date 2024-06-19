package apiserver

import "time"

// Config ...
type Config struct {
	BindAddr    string        `toml:"bind_addr"`
	LogLevel    string        `toml:"log_level"`
	DatabaseURL string        `toml:"database_url"`
	SessionKey  string        `toml:"session_key"`
	Timeout     time.Duration `toml:"timeout"`
	IdleTimout  time.Duration `toml:"idle_timeout"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		LogLevel:   "debug",
		Timeout:    14 * time.Second,
		IdleTimout: 60 * time.Second,
	}
}
