package posts

import (
	"context"

	"github.com/NXRts/fsatcampus/internal/middleware"
	"github.com/NXRts/fsatcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, UserId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
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
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpsertUserActivity)
}
