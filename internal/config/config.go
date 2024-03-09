package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Initialize initializes the viper instance. It will only run once.
func Initialize() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	filePath := filepath.Join(homedir, ".acli")

	viper.SetConfigName(".acli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(homedir)
	err = viper.ReadInConfig()
	if err == nil {
		return nil
	}
	_, ok := err.(viper.ConfigFileNotFoundError)
	if !ok {
		// Config file was found but another error was produced
		return fmt.Errorf("failed to get acli config file: %w", err)

	}
	return writeConfigFile(filePath)
}

// writeConfigFile creates a new acli config file in the user's home directory.
func writeConfigFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to create acli config file: %w", err)
	}
	file.Close()
	return nil
}

// Get reads a config string from the acli config file.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStringSlice reads a config string slice from the acli config file.
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// Set writes a config value to the acli config file.
func Set(key string, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error writing viper config file: %w", err))
	}
}
