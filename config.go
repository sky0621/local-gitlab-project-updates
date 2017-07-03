package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	GitlabApiUrl string
	PrivateToken string
	GitlabSshUrl string
	Outdir       string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		GitlabApiUrl: viper.GetString("gitlabApiUrl"),
		PrivateToken: viper.GetString("privateToken"),
		GitlabSshUrl: viper.GetString("gitlabSshUrl"),
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
