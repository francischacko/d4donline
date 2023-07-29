package configs

import (
	"log"

	"github.com/spf13/viper"
)

type EnvConfigs struct {
	LocalServerPort string `mapstructure:"port"`
	DbConnect       string `mapstructure:"dsn"`
}

var EnvConf *EnvConfigs

func InitEnvConfigs() {
	EnvConf = LoadEnvVariables()
}

func LoadEnvVariables() (configs *EnvConfigs) {
	viper.AddConfigPath(".")

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading env env variables")
	}

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal("Error while unmarshalling loaded variables into struct")
	}
	return
}
