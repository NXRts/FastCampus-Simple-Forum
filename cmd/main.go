package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yeremiaaryo96/fastcampus/internal/configs"
	"github.com/yeremiaaryo96/fastcampus/internal/handlers/memberships"
	"github.com/yeremiaaryo96/fastcampus/internal/handlers/posts"
	membershipRepo "github.com/yeremiaaryo96/fastcampus/internal/repository/memberships"
	postRepo "github.com/yeremiaaryo96/fastcampus/internal/repository/posts"
	membershipSvc "github.com/yeremiaaryo96/fastcampus/internal/service/memberships"
	postSvc "github.com/yeremiaaryo96/fastcampus/internal/service/posts"
	"github.com/yeremiaaryo96/fastcampus/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
