package main

import (
	"log"

	config "github.com/NXRts/fsatcampus/internal/configs"
	"github.com/NXRts/fsatcampus/internal/handlers/memberships"
	"github.com/NXRts/fsatcampus/internal/handlers/posts"
	membershipsRepo "github.com/NXRts/fsatcampus/internal/repository/memberships"
	postRepo "github.com/NXRts/fsatcampus/internal/repository/posts"
	membershipsSvc "github.com/NXRts/fsatcampus/internal/service/memberships"
	postSvc "github.com/NXRts/fsatcampus/internal/service/posts"
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
	// log.Println("Config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal Inisialisasi Databases", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipsRepo := membershipsRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipsService := membershipsSvc.NewService(cfg, membershipsRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipsHandler := memberships.NewHandler(r, membershipsService)
	membershipsHandler.RegisterRoutes()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
