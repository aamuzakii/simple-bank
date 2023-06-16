package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DBDriver"`
	DBSource      string `mapstructure:"DBSource"`
	ServerAddress string `mapstructure:"ServerAddress"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
