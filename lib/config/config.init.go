package config

import (
	"github.com/spf13/viper"
)

func Init(configName, path string) (*Config, error) {
	cfg := &Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	return cfg, err
}
