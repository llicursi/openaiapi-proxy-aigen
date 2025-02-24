package config

import (
	"fmt"
	"os"
	"path/filepath"

	"inspect-proxy/internal/handlers"

	"gopkg.in/yaml.v3"
)

// LoadConfig loads configuration from a YAML file based on the environment
func LoadConfig() (handlers.Config, error) {
	var config handlers.Config

	// Get environment, default to development
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Construct config filename
	filename := fmt.Sprintf("config.%s.yaml", env)

	// Look for config file in common locations
	configPaths := []string{
		filename,
		filepath.Join("config", filename),
		filepath.Join("/etc/inspect-proxy", filename),
	}

	var configFile string
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			configFile = path
			break
		}
	}

	if configFile == "" {
		return config, fmt.Errorf("no config file found for environment: %s", env)
	}

	// Read config file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse YAML
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config file: %w", err)
	}

	// Allow override of forward URL through environment variable
	if forwardURL := os.Getenv("FORWARD_URL"); forwardURL != "" {
		config.Server.ForwardURL = forwardURL
	}

	return config, nil
}
