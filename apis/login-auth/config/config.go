package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct{
  API APIConfig
  DB  DBConfig
}

type APIConfig struct{
  Port string 
}

type DBConfig struct{
  Host string
  Port string
  User string
  Pass string
  DB   string
}

func init()  {
  viper.SetDefault("api.port", "8080")
  viper.SetDefault("database.host", "localhost")
  viper.SetDefault("database.port", "5432")
}

func Load() error {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("Config error file: %w", err))
  }

  cfg = new(Config)

  cfg.API = APIConfig{
    Port: viper.GetString("api.port"),
  }

  cfg.DB = DBConfig{
    Host: viper.GetString("database.host"),
    Port: viper.GetString("database.port"),
    User: viper.GetString("database.user"),
    Pass: viper.GetString("database.pass"),
    DB:   viper.GetString("database.name"),
  }

  return nil
}


func GetDB() DBConfig {
  return cfg.DB
}

func GetServerPort() string {
  return cfg.API.Port
}
