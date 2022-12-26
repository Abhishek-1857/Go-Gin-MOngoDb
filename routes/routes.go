package routes


import (
	"github.com/gin-gonic/gin"
     "oracleservice/controller"
)

func Routes(router *gin.Engine)  {
   router.POST("/add",controllers.AddInfo())
   router.GET("/get",controllers.GetInfo())
   router.GET("get/ipfshash",controllers.GetIpfsHashes())
}