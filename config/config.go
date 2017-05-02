package config

import (
	"encoding/json"
    "strings"
	"io/ioutil"
	"os"
    "fmt"
    "gopkg.in/yaml.v2"
)

type DatabaseSettings struct {
    Path string `json:"path" yaml:"path"`
    PostgresUrl string `json:"postgres_url" yaml:"postgres_url"`
}

type StorageSettings struct {
    Path string `json:"path" yaml:"path"`
    Bucket string `json:"bucket" yaml:"bucket"`
    Credentials string `json:"credentials" yaml:"credentials"`
}

type Config struct {
    Port string `json:"port" yaml:"port"`
    Database string `json:"database" yaml:"database"`
    DatabaseSettings DatabaseSettings `json:"database_settings" yaml:"database_settings"`
    StorageBackend string `json:"storage_backend" yaml:"storage_backend"`
    StorageSettings StorageSettings `json:"storage_settings" yaml:"storage_settings"`
}

func NewConfig(env []string) (*Config, error) {
    result := &Config{ }
    replaceMissingValuesWithDefaults(result)
    return processEnvironmentOverrides(result, env), nil
}

func replaceMissingValuesWithDefaults(config *Config) {
    if config.Port == "" {
        config.Port = "7770"
    }
    if config.Database == "" {
        config.Database = "sqlite"
    }
    if config.DatabaseSettings.Path == "" && config.Database == "sqlite" {
        config.DatabaseSettings.Path = "/var/lib/escape/registry.db"
    }
    if config.StorageBackend == "" {
        config.StorageBackend = "local"
    }
    if config.StorageSettings.Path == "" && config.StorageBackend == "local" {
        config.StorageSettings.Path = "/var/lib/escape/releases"
    }
}

func LoadConfig(file string, env []string) (*Config, error) {
	var config Config

	if !PathExists(file) {
		return nil, fmt.Errorf("The referenced configuration file '%s' could not be found", file)
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
        return nil, fmt.Errorf("Error reading configuration file '%s': %s", file, err.Error())
	}

    if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
        if err = yaml.Unmarshal(b, &config); err != nil {
            return nil, fmt.Errorf("Could not parse YAML in configuration file '%s': %s", file, err.Error())
        }
    } else {
        if err = json.Unmarshal(b, &config); err != nil {
            return nil, fmt.Errorf("Could not parse JSON in configuration file '%s': %s", file, err.Error())
        }
    }
    replaceMissingValuesWithDefaults(&config)
    return processEnvironmentOverrides(&config, env), nil
}

func processEnvironmentOverrides(config *Config, env []string) *Config {
    for _, e := range env {
        parts := strings.SplitN(e, "=", 2)
        key := parts[0]
        value := parts[1]
        if key == "PORT" {
            config.Port = value
        } else if key == "DATABASE" {
            config.Database = value
        } else if key == "DATABASE_SETTINGS_PATH" {
            config.DatabaseSettings.Path = value
        } else if key == "DATABASE_SETTINGS_POSTGRES_URL" {
            config.DatabaseSettings.PostgresUrl = value
        } else if key == "STORAGE_BACKEND" {
            config.StorageBackend = value
        } else if key == "STORAGE_SETTINGS_PATH" {
            config.StorageSettings.Path = value
        } else if key == "STORAGE_SETTINGS_BUCKET" {
            config.StorageSettings.Bucket = value
        } else if key == "STORAGE_SETTINGS_CREDENTIALS" {
            config.StorageSettings.Credentials = value
        }
    }
    return config
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}