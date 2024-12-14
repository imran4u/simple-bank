package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // it can be json, xml etc other file formate

	viper.AutomaticEnv() // setup to override env values if exist

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return // this will return config and err , it is short form to write only return
}
