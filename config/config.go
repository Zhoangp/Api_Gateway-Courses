package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Services ServicesConfig
	Email    EmailConfig
}
type EmailConfig struct {
	AppEmail    string
	AppPassword string
}
type ServicesConfig struct {
	Port                  string
	FileUrl string
	AuthUrl               string
	CourseUrl             string
	CategoryUrl           string
	TopicUrl              string
	Secret                string
	UserUrl               string
	AccessTokenExpiredIn  int
	RefreshTokenExpiredIn int
}

func LoadConfig(fileName string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cf Config
	if err := v.Unmarshal(&cf); err != nil {
		return nil, err
	}

	return &cf, nil
}
