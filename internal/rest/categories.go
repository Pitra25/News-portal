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

	c.JSON(
		http.StatusOK,
		NewCategoryArr(categories),
	)
}
