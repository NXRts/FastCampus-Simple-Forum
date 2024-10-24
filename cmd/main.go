package main

import (
	"log"

	config "github.com/NXRts/fsatcampus/internal/configs"
	"github.com/NXRts/fsatcampus/internal/handlers/memberships"
	membershipsRepo "github.com/NXRts/fsatcampus/internal/repository/memberships"
	membershipsSvc "github.com/NXRts/fsatcampus/internal/service/memberships"
	"github.com/NXRts/fsatcampus/pkg/internalsql"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal Inisialisasi Databases", err)
	}

	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipsService := membershipsSvc.NewService(membershipsRepo)

	membershipsHandler := memberships.NewHandler(r, membershipsService)
	membershipsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
