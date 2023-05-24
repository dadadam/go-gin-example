package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port       int    `mapstructure:"PORT"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int    `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
}

var config *Config

func (c *Config) GetDsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.DbHost, c.DbUser, c.DbPassword, c.DbName, c.DbPort,
	)
}

func (c *Config) GetDbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName,
	)
}

func setDefaults(loader *viper.Viper) {
	loader.SetDefault("PORT", 8080)
	loader.SetDefault("LOG_LEVEL", "INFO")
	loader.SetDefault("DB_HOST", "localhost")
	loader.SetDefault("DB_PORT", 5432)
	loader.SetDefault("DB_NAME", "app_db")
	loader.SetDefault("DB_USER", "app_user")
	loader.SetDefault("DB_PASSWORD", "")
}

func Init() {
	var err error

	loader := viper.New()
	loader.SetConfigType("env")
	loader.SetConfigName(".env")
	loader.AddConfigPath(".")

	setDefaults(loader)

	loader.AutomaticEnv()
	viper.WatchConfig()

	err = loader.ReadInConfig()
	if err != nil {
		log.Warn(err)
	}

	err = loader.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}
