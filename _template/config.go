package config

import "github.com/spf13/viper"

// Config ...
type Config struct {
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
