package main

import (
	"github.com/binganao/dailyCVE/lib"
	"github.com/binganao/dailyCVE/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	lib.InitConf()
	db := lib.InitDB()
	defer db.Close()
	c := lib.SetCron()
	c.Start()
	defer c.Stop()
	r := gin.Default()
	r.GET("/api/cve/list", func(ctx *gin.Context) {
		date := ctx.Query("date")
		ctx.JSON(http.StatusOK, model.Resp{Code: http.StatusOK, CVEs: lib.QueryCVE(date)})
	})
	r.GET("/api/cve/daily", func(ctx *gin.Context) {
		// 获取到的是昨天的内容
		date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		ctx.JSON(http.StatusOK, model.Resp{Code: http.StatusOK, CVEs: lib.QueryCVE(date)})
	})
	r.Run("0.0.0.0:" + viper.GetString("http.port"))
}
