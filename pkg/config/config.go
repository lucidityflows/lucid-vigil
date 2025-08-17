package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string `mapstructure:"log_level"`
	APIPort string `mapstructure:"api_port"`
	Monitors []MonitorConfig `mapstructure:"monitors"`
	Actions  ActionsConfig   `mapstructure:"actions"`
}

type MonitorConfig struct {
	Name string `mapstructure:"name"`
	Enabled bool `mapstructure:"enabled"`
	Interval string `mapstructure:"interval"`
	// Add more monitor-specific configurations here
	Actions []string `mapstructure:"actions"` // Actions to trigger for this monitor
}

type ActionsConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Add global action-specific configurations here
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config") // config.yaml
	v.SetConfigType("yaml")
	v.AddConfigPath(".") // Search in current directory
	v.AddConfigPath("/etc/sentinel/") // Search in /etc/sentinel/

	// Set default values
	v.SetDefault("log_level", "info")
	v.SetDefault("api_port", "8080") // Default API port
	v.SetDefault("actions.enabled", false) // Actions disabled by default

	// Read environment variables
	v.SetEnvPrefix("SENTINEL") // Look for SENTINEL_ prefix
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace dots with underscores for nested keys
	v.AutomaticEnv() // Automatically bind environment variables to config keys

	// Read configuration file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, using defaults and environment variables.")
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
