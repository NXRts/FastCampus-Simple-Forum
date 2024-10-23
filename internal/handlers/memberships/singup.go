package memberships

import (
	"net/http"

	"github.com/NXRts/fsatcampus/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SingUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SingUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := h.membershipSvc.SingUp(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
