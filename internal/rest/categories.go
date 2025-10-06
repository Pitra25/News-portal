package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Router) GetAllCategories(c *gin.Context) {

	categories, err := h.newsportal.GetAllCategory()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var category []Category
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
