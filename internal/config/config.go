package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port           string   `yaml:"port"`
		TrustedProxies []string `yaml:"trusted_proxies"`
		APIKey         []string `yaml:"apiKey"`
	} `yaml:"server"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	config.Server.Port = "8080"
	config.Server.TrustedProxies = []string{}
	config.Server.APIKey = []string{} // Default to empty list, accepting any key

	if configPath == "" {
		return config, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfigPath() string {
	if _, err := os.Stat("config.yaml"); err == nil {
		return "config.yaml"
	}

	home, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(home, ".hrng-rpc", "config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}

	return ""
}
