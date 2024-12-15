package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev") // Nama file tanpa ekstensi
	viper.SetConfigType("env") // Tipe file

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	err := viper.Unmarshal(&config)
	return config, err
}
