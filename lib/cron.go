package lib

import (
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func SetCron() *cron.Cron {
	GetCVE()
	c := cron.New()
	c.AddFunc(viper.GetString("cron.spec"), GetCVE)
	return c
}
