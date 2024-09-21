package router

import (
	//"simulation/src/apisimulation/controller/router"

	user "simulation/src/apisimulation/controller/user"
	"simulation/src/apisimulation/controller/utils"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
// Create path url 
router := gin.New()

router.Use(requestid.New())

return router
}

func InitRouter(router *gin.Engine)  {
	v1public := router.Group("/api/v1/public")
	{
		v1public.POST("/login",user.Login)
		v1public.POST("/register",user.Register)
		v1public.POST("getUser/id")
		v1public.GET("getAllUser")
		v1public.POST("deleteUser/id")
		v1public.GET("/tes",user.Tes)
	}
	v1private := router.Group("/api/v1/private")
	v1private.Use(utils.AuthMiddleware())
	{
		v1private.GET("/tes",user.GetUserDetails)
	}
}

