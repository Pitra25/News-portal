package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Router) GetAllTags(c *gin.Context) {
	tags, err := h.newsportal.GetAllTag()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": tags,
		},
	)
}
