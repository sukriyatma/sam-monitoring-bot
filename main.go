package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {ctx.IndentedJSON(200, "OK")})
	router.GET("/getbots", getBotsActivity)	
	router.GET("/login", validateUser)
	router.GET("/findmonitors", findMonitors)
	router.GET("/findbotsbystatus", findBotsbyStatus)

	router.POST("/removemonitor", removeMonitor)
	router.POST("/insertbot", postBotsActivity)
	router.POST("/insertuser", postUser)

	router.Run("localhost:8080")
}


