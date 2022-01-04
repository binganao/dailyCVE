package lib

import (
	"github.com/spf13/viper"
	"log"
)

func InitConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// gin setting
	viper.SetDefault("http.port", 80)

	// cron setting
	viper.SetDefault("cron.spec", "0 30 12 * * ?")

	// mysql setting
	viper.SetDefault("mysql.host", "localhost")
	viper.SetDefault("mysql.port", "3306")
	viper.SetDefault("mysql.database", "daily_cve")
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "root")
	viper.SetDefault("mysql.charset", "utf8mb4")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	//viper.WriteConfigAs("config.toml")
}
