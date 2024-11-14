package posts

import (
	"net/http"

	"github.com/NXRts/fsatcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetInt64("userID")
	err := h.postSvc.CreatePost(ctx, userId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

// func (h *Handler) CreatePost(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	var request posts.CreatePostRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	userId, exists := c.Get("userId")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "user not authenticated",
// 		})
// 		return
// 	}

// 	userIdInt64, ok := userId.(int64)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "user ID is not valid",
// 		})
// 		return
// 	}

// 	err := h.postSvc.CreatePost(ctx, userIdInt64, request)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.Status(http.StatusCreated)
// }
