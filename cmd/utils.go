package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

// getConfigDir returns the Kalco configuration directory
func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".kalco")
	return configDir, nil
}
