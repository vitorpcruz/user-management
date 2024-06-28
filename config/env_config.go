package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type envConfig struct {
	appPort  int64  `mapstructure:"APP_PORT"`
	mySqlDSN string `mapstructure:"MYSQL_DSN"`
}

var (
	lock   = &sync.Mutex{}
	config *envConfig
)

func LoadEnvConfig(path string) {
	if config != nil {
		return
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	lock.Lock()
	defer lock.Unlock()
	config = &envConfig{
		appPort:  viper.GetInt64("APP_PORT"),
		mySqlDSN: viper.GetString("MYSQL_DSN"),
	}
}

func AppPort() int64 {
	return config.appPort
}

func MySqlDSN() string {
	return config.mySqlDSN
}
