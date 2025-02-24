package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary directory for test config files
	tmpDir, err := os.MkdirTemp("", "config-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test config files for both test and development environments
	testConfig := `
server:
  forward_url: "http://test-api.com"
`
	// Write test environment config
	err = os.WriteFile(filepath.Join(tmpDir, "config.test.yaml"), []byte(testConfig), 0644)
	require.NoError(t, err)

	// Write development environment config
	err = os.WriteFile(filepath.Join(tmpDir, "config.development.yaml"), []byte(testConfig), 0644)
	require.NoError(t, err)

	tests := []struct {
		name        string
		env         string
		forwardURL  string
		configDir   string
		expectError bool
	}{
		{
			name:      "valid config with ENV",
			env:       "test",
			configDir: tmpDir,
		},
		{
			name:      "default to development",
			configDir: tmpDir,
		},
		{
			name:        "non-existent config",
			env:         "nonexistent",
			expectError: true,
		},
		{
			name:       "override forward URL",
			env:        "test",
			configDir:  tmpDir,
			forwardURL: "http://override-url.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment
			if tt.env != "" {
				os.Setenv("ENV", tt.env)
				defer os.Unsetenv("ENV")
			} else {
				os.Unsetenv("ENV")
			}

			if tt.forwardURL != "" {
				os.Setenv("FORWARD_URL", tt.forwardURL)
				defer os.Unsetenv("FORWARD_URL")
			} else {
				os.Unsetenv("FORWARD_URL")
			}

			// Change working directory if specified
			if tt.configDir != "" {
				currentDir, err := os.Getwd()
				require.NoError(t, err)
				err = os.Chdir(tt.configDir)
				require.NoError(t, err)
				defer os.Chdir(currentDir)
			}

			// Load config
			config, err := LoadConfig()

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			// Verify config values
			if tt.forwardURL != "" {
				assert.Equal(t, tt.forwardURL, config.Server.ForwardURL)
			} else {
				assert.Equal(t, "http://test-api.com", config.Server.ForwardURL)
			}
		})
	}
}
