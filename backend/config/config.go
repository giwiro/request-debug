package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Server      struct {
		Port     string
		Address  string
		BasePath string `mapstructure:"base_path"`
	}
	Logger struct {
		Level string
	}
	Database struct {
		Uri    string
		DBName string
	}
	App struct {
		Name string
	}
}

var Conf Config

func ReadConfig(path string) error {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}

	fmt.Printf("Loading config from: %s", path)

	conf, _ := json.MarshalIndent(Conf, "", "\t")
	fmt.Printf("%s", conf)

	return nil
}
