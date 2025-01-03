package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type DataBase struct {
	DNS string `mapstructure:"DNS"`
}
type ServerInfo struct {
	Port   string  `mapstructure:"port"`
	Name   *string `mapstructure:"name"`
	Static string  `mapstructure:"static"`
	Html   string  `mapstructure:"html"`
}
type FmConfig struct {
	Server      ServerInfo `mapstructure:"server"`
	Database    DataBase   `mapstructure:"database"`
	Version     *string    `mapstructure:"version"`
	Environment *string    `mapstructure:"environment"`
}

var AppConfig FmConfig

func Init() error {

	exePath, err := os.Executable()

	if err != nil {

		log.Fatalf("Error reading executeable file path, %s", err)
		return err
	}

	return InitPath(exePath)
}

func InitPath(exePath string) error {

	viper.SetEnvPrefix("APP")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// auto read environment variable values
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(exePath))

	configFilePath := filepath.Join(filepath.Dir(exePath), "config.yaml")

	// Check if the config file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist at path %s", configFilePath)
		return err
	} else {
		viper.SetConfigFile(configFilePath)
	}

	log.Printf("Loaded config from %s", configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}
	// print all settings in viper
	log.Println(viper.AllSettings())

	if err := viper.Unmarshal(&AppConfig, func(m *mapstructure.DecoderConfig) {
		m.TagName = "mapstructure"
	}); err != nil {
		log.Fatalf("Failed to decode config: %v", err)
		return err
	}

	return nil
}
