package router

import (
	//"simulation/src/apisimulation/controller/router"

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
		v1public.POST("login")
		v1public.POST("register")
		v1public.POST("getUser/id")
		v1public.GET("getAllUser")
		v1public.POST("deleteUser/id")
	}
}