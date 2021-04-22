package util

import (
	"github.com/spf13/viper"
)


type Config struct {
	Source     string `mapstructure:"SOURCE"`
	SourceType string `mapstructure:"SOURCE_TYPE"`
	Host       string `mapstructure:"HOST"`
}

func ReadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("settings")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}