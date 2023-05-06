package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {ctx.Status(200)})
	router.GET("/monitoringbot/getbots", getBotsActivity)	
	router.GET("/monitoringbot/login", validateUser)
	router.GET("/monitoringbot/findmonitors", findMonitors)
	router.GET("/monitoringbot/findbotsbystatus", findBotsbyStatus)

	router.POST("/monitoringbot/removemonitor", removeMonitor)
	router.POST("/monitoringbot/insertbot", postBotsActivity)
	router.POST("/monitoringbot/insertuser", postUser)

	router.Run("127.0.0.1:3000")
}


