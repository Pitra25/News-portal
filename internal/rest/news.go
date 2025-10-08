package rest

import (
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
	filter := params.NewFilter()

	news, err := h.manager.GetNewsByFilters(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		NewNewsArr(news),
	)
}

func (h *Router) GetNewsById(c *gin.Context) {
	newsIdStr, _ := c.Params.Get("id")

	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	news, err := h.manager.GetNewsById(newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	filter := params.NewFilter()

	shortNews, err := h.manager.GetNewsByFilters(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		NewShortNews(shortNews),
	)
}

func (h *Router) GetNewsCount(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filter := params.NewFilter()

	count, err := h.manager.GetNewsCount(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		count,
	)
}
