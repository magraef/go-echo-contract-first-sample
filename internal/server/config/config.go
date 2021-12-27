package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port       uint16 `default:"8080"`
	ApiBaseUri string `split_words:"true" default:"api"`
	Debug      bool   `default:"false"`
	JsonLog    bool   `split_words:"true" default:"false"`
}


func Provide() *Config {
	var c Config
	err := envconfig.Process("app", &c)
	if err != nil {
		logrus.Panicf(err.Error())
	}
	return &c
}
