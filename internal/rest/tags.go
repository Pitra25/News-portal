package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Router) GetAllTags(c *gin.Context) {
	tagsArr, err := h.newsportal.GetAllTag()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tags := make([]Tag, len(tagsArr))
	for _, v := range tagsArr {
		tags = append(tags, NewTag(v))
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": tags,
		},
	)
}
