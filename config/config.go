package config

import (
	"sync"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Port   int    `toml:"port" mapstructure:"port"`
		JWTKey string `toml:"jwtkey" mapstructure:"jwtkey"`
	} `toml:"app" mapstructure:"app"`
	Database struct {
		Driver      string `toml:"driver" mapstructure:"driver"`
		DB_HOST     string `toml:"DB_HOST" mapstructure:"dbhost"`
		DB_NAME     string `toml:"DB_NAME" mapstructure:"dbname"`
		DB_PORT     string `toml:"DB_PORT" mapstructure:"dbport"`
		DB_USER     string `toml:"DB_USER" mapstructure:"user"`
		DB_PASSWORD string `toml:"DB_PASSWORD" mapstructure:"password"`
	} `toml:"database"`
	Log struct {
		Driver string `toml:"driver" mapstructure:"driver"`
	} `toml:"log"`
	OpenApi struct {
		MusixMatch    string `toml:"musixmatch" mapstructure:"musixmatch"`
		MusixMatchUrl string `toml:"musixmatchurl" mapstructure:"musixmatchurl"`
	} `toml:"openapi"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.App.Port = 5006

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("error when load config file", err)
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("error when parse config file", err)
		return &defaultConfig
	}
	return &finalConfig
}
