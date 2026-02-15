package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds application configuration.
type Config struct {
	Theme    string `mapstructure:"theme"`
	LogLevel string `mapstructure:"log_level"`
}

// Load reads config from ~/.avro/config.yaml or defaults.
func Load() *Config {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".avro")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	viper.SetDefault("theme", "default")
	viper.SetDefault("log_level", "info")

	viper.SetEnvPrefix("AVRO")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig() // ignore missing config

	cfg := &Config{}
	_ = viper.Unmarshal(cfg)
	return cfg
}
