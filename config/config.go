package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type RssConfig struct {
	Ria  string
	Tass string
}

type Config struct {
	Rss RssConfig
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("cannot map config: %w", err)
	}

	return &cfg, nil
}
