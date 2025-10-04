package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Router) GetAllCategories(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filter := params.NewFilter()

	categories, err := h.newsportal.GetAllCategory(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	category := make([]Category, len(categories))
	for _, v := range categories {
		category = append(category, NewCategory(v))
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": category,
		},
	)
}
