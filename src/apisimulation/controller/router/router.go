package router

import (
	//"simulation/src/apisimulation/controller/router"

	user "simulation/src/apisimulation/controller/user"
	product "simulation/src/apisimulation/controller/product"
	"simulation/src/apisimulation/controller/utils"
	simulasi "simulation/src/apisimulation/controller/simulation"

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
		v1public.GET("/lob",product.Lob)
		v1public.POST("/lob/id",product.LobById)

		//v1public.GET("/tes",user.Tes)
	}
	v1private := router.Group("/api/v1/private")
	v1private.Use(utils.AuthMiddleware())
	{
		//v1private.GET("/",product.GetUserInfo)

		// lob MYB
		v1private.GET("/category/newmotorcycle",product.MasterCatNewMotorcycle)
		v1private.POST("/product/newmotorcycle",product.ProductNewMotorcycle)
		v1private.POST("/price/newmotorcycle",product.PriceNewMotorcycle)

		// lob MB
		v1private.GET("/category/usedmotorcycle",product.MasterCatIUsedMotorcycle)
		v1private.POST("/product/usedmotorcycle",product.ProductUsedMotorcycle)
		v1private.POST("/price/usedmotorcycle",product.PriceUsedMotorcycle)

		// lob car
		v1private.GET("/category/car",product.MasterCatCar)
		v1private.POST("/product/car",product.ProductCar)
		v1private.POST("/price/car",product.PriceCar)

		// lob Mp
		v1private.GET("/category/multiproduct",product.MasterMultiproduct)
		v1private.POST("/product/multiproduct",product.ProductMultiproduct)
		v1private.POST("/price/multiproduct",product.PriceMultiproduct)

		// simulasi
		v1private.POST("/simulation",simulasi.Simulation)


	}
}

