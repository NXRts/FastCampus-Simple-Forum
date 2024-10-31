package posts

import (
	"context"

	"github.com/NXRts/fsatcampus/internal/middleware"
	"github.com/NXRts/fsatcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, UserId int64, req posts.CreatePostRequest) error
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create_post", h.CreatePost)
}
