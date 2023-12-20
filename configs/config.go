package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var configuration *Config

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBTimezone   string `mapstructure:"DB_TIMEZONE"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
	JWTSecret    string `mapstructure:"SERVER_PORT"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

func Configure(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&configuration)
	if err != nil {
		return nil, err
	}
	configuration.TokenAuth = jwtauth.New("HS256", []byte(configuration.JWTSecret), nil)
	return configuration, nil
}
