package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Router) GetAllNews(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	list, err := h.manager.GetNewsByFilters(c.Request.Context(), params.NewFilter())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if len(list) == 0 {
		newErrorResponse(c, http.StatusNoContent, errors.New("no list found"))
		return
	}

	c.JSON(
		http.StatusOK,
		NewNewsList(list),
	)
}

func (h *Router) GetNewsById(c *gin.Context) {
	newsIdStr, _ := c.Params.Get("id")

	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	news, err := h.manager.GetNewsById(c.Request.Context(), newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK,
		NewNews(news),
	)
}

func (h *Router) GetAllShortNews(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	list, err := h.manager.GetNewsByFilters(c.Request.Context(), params.NewFilter())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if len(list) == 0 {
		newErrorResponse(c, http.StatusNoContent, errors.New("no news found"))
		return
	}

	c.JSON(
		http.StatusOK,
		NewNewsSummaries(list),
	)
}

func (h *Router) GetNewsCount(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	count, err := h.manager.GetNewsCount(c.Request.Context(), params.NewFilter())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK,
		count,
	)
}

func (h *Router) GetAllCategories(c *gin.Context) {

	categories, err := h.manager.GetAllCategory(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK,
		NewCategories(categories),
	)
}

func (h *Router) GetAllTags(c *gin.Context) {
	tags, err := h.manager.GetAllTag(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK,
		NewTags(tags),
	)
}
