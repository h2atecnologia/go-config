// Package config provides super simple JSON configuration for command-line programs.
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Load configuration from path, into the struct pointer provided.
func Load(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

// Save saves configuration to path.
func Save(path string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, 0600)
}

// LoadHome loads configuration from path relative to the user home directory.
func LoadHome(path string, v interface{}) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return Load(filepath.Join(home, path), v)
}

// SaveHome saves configuration to path relative to the user home directory.
func SaveHome(path string, v interface{}) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return Save(filepath.Join(home, path), v)
}
