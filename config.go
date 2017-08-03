package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Scheme       string
	Host         string
	User         string
	Pass         string
	Branch       string
	PrivateToken string
	Outdir       string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Scheme:       viper.GetString("scheme"),
		Host:         viper.GetString("host"),
		User:         viper.GetString("user"),
		Pass:         viper.GetString("pass"),
		Branch:       viper.GetString("branch"),
		PrivateToken: viper.GetString("privateToken"),
		Outdir:       viper.GetString("outdir"),
	}
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}

// Host4GitCommand ...
func (c *Config) Host4GitCommand(pathWithNamespace string) string {
	return fmt.Sprintf("%s/%s.git", c.GitlabSshUrl, pathWithNamespace)
}
