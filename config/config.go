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
	} `toml:"app"`
	Database struct {
		Driver      string `toml:"driver" mapstructure:"driver"`
		DBURL       string `toml:"dburl" mapstructure:"dburl"`
		DB_HOST     string `toml:"dburl" mapstructure:"dbhost"`
		DB_NAME     string `toml:"dburl" mapstructure:"dbname"`
		DB_PORT     string `toml:"dburl" mapstructure:"dbport"`
		DB_USER     string `toml:"dburl" mapstructure:"user"`
		DB_PASSWORD string `toml:"dburl" mapstructure:"password"`
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

	viper.SetConfigName("app")
	viper.SetConfigType("env")
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
