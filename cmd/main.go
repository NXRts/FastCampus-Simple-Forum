package main

import (
	"log"

	config "github.com/NXRts/fsatcampus/internal/configs"
	"github.com/NXRts/fsatcampus/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *config.Config
	)

	err := config.Init(
		config.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal Inisialisasi Config", err)
	}

	cfg = config.Get()
	log.Println("Config", cfg)

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
