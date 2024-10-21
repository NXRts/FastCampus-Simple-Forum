package main

import (
	"github.com/NXRts/fsatcampus/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ganti NawHandler dengan NewHandler
	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoutes()

	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
