package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeremiaaryo96/fastcampus/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.membershipSvc.SignUp(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
