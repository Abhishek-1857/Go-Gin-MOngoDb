package main

import (
	"oracleservice/config"
	"oracleservice/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	//connect database
	config.ConnectDB()

	//routes4
    routes.Routes(router)

	router.Run("localhost:6000")
}
