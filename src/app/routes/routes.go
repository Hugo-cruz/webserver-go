package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	handler "webserver/src/app/handler"
	"webserver/src/app/usecase"
	dep "webserver/src/dependencies"
)

func Router(dep dep.Dependencies) {
	useCase := usecase.NewUseCase(dep.DB)
	deviceHandler := handler.NewDevicesHandler(useCase)

	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "home")
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	deviceGroup := router.Group("/device")

	deviceGroup.GET("/initialize", deviceHandler.InitializeDB)
	deviceGroup.GET("/:id", deviceHandler.GetDeviceByID)
	deviceGroup.GET("/search", deviceHandler.SearchByBrand)
	deviceGroup.GET("/list", deviceHandler.ListDevices)
	deviceGroup.POST("/create", deviceHandler.CreateDevice)
	deviceGroup.DELETE("/:id", deviceHandler.DeleteDevice)
	deviceGroup.PUT("/:id", deviceHandler.UpdateDevice)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
