package lib

import (
	"github.com/spf13/viper"
	"log"
)

func InitConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("http.port", 80)
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("cron.spec", "0 30 12 * * ?")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	viper.WriteConfigAs("config.toml")
}
