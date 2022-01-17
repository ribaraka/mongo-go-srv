package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mongo        MongoConfiguration
	StaticAssets string `mapstructure:"staticAssets"`
	Mail         MailConfig
}

type MongoConfiguration struct {
	ServerHost string `mapstructure:"server"`
	Database   string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
	Credentials MongoCred `mapstructure:"credentials"`
}

type MongoCred struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type MailConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Sender       string `mapstructure:"sender"`
	AuthEmail    string `mapstructure:"user"`
	AuthPassword string `mapstructure:"password"`
}

func LoadConfig(file string) (config Config, err error) {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()
	err = viper.BindEnv("mail.sender", "MAIL_SENDER")
	if err != nil {
		return
	}

	err = viper.BindEnv("mail.user", "MAIL_USER")
	if err != nil {
		return
	}

	err = viper.BindEnv("mail.password", "MAIL_PASSWORD")
	if err != nil {
		return
	}

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
