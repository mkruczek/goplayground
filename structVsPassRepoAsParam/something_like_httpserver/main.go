package something_like_httpserver

import (
	"context"
	"temporary/gin"
)

func mainStruct() {
	deviceRepository := DeviceRepository(nil)

	auth := DeviceAuthentication{
		DeviceRepo: deviceRepository,
	}

	r := gin.Gin{}
	r.GET("/ping",
		auth.AuthenticateMiddleware(),
		func(c context.Context) {
			//some code
		})

	r.Run()
}

func mainFunc() {
	deviceRepository := DeviceRepository(nil)

	r := gin.Gin{}
	r.GET("/ping",
		AuthenticateMiddleware(deviceRepository),
		func(c context.Context) {
			//some code
		})

	r.Run()
}
