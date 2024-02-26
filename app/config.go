package app

import (
	"github.com/go-errors/errors"
	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP          HTTP          `mapstructure:"http"`
		Authorization Authorization `mapstructure:"authorization"`
		Postgres      Postgres      `mapstructure:"postgres"`
		Cache         Cache         `mapstructure:"cache"`
		Storage       Storage       `mapstructure:"storage"`
	}
	HTTP struct {
		Port uint16
	}
	Authorization struct {
		JWTSigningKey    string `mapstructure:"jwtSigningKey"`
		PasswordHashSalt string `mapstructure:"passwordHashSalt"`
	}
	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslMode"`
	}
	Cache struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	}
	Storage struct {
		Endpoint          string `mapstructure:"endpoint"`
		AccessKeyId       string `mapstructure:"accessKeyId"`
		SecretAccessKeyId string `mapstructure:"secretAccessKeyId"`
	}
)

// loadConfig - initializing app config from yaml file
func loadConfig(configFile string) (appCfg Config, err error) {
	viper.SetConfigFile(configFile)
	if err = viper.ReadInConfig(); err != nil {
		err = errors.New(err)
		return
	}
	if err = viper.Unmarshal(&appCfg); err != nil {
		err = errors.New(err)
		return
	}

	return
}
