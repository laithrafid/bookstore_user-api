package config_utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	/// ConfigVarName:Type:MaptoConfigVarNameindotenvfile
	ServerAddress string `mapstructure:"ServerAddress"`
	Username      string `mapstructure:"mysqlUsersUsername"`
	Password      string `mapstructure:"mysqlUsersPassword"`
	Host          string `mapstructure:"mysqlUsersHost"`
	Schema        string `mapstructure:"mysqlUsersSchema"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
