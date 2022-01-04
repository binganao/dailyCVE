package lib

import (
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"log"
)

func SetCron()  {
	c := cron.New()

	if viper.Get("cron.spec") == "" {
		log.Println(viper.Get("cron.spec"))
	}
	c.AddFunc("0 30 12 * * ?", GetCVE)
}
