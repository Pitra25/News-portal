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

	newsDB, err := h.newsportal.GetNewsByFilters(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	news := make([]News, len(newsDB))
	for _, v := range newsDB {
		tags := make([]Tag, len(v.Tags))
		for _, tag := range v.Tags {
			tags = append(tags, NewTag(tag))
		}
		news = append(news, NewNews(v, tags))
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": news,
		},
	)
}

func (h *Router) GetNewsById(c *gin.Context) {
	newsIdStr, _ := c.Params.Get("id")

	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newsArr, err := h.newsportal.GetNewsById(newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tags := make([]Tag, len(newsArr.Tags))
	for _, tag := range tags {
		tags = append(tags, tag)
	}

	news := NewNews(newsArr, tags)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": news,
		},
	)
}

func (h *Router) GetAllShortNews(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filter := params.NewFilter()

	shortNewsArr, err := h.newsportal.GetALlShortNewsByFilters(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	shortNews := make([]ShortNews, len(shortNewsArr))
	for i, v := range shortNewsArr {
		tags := make([]Tag, len(v.Tags))
		for _, tag := range tags {
			tags = append(tags, tag)
		}

		shortNews = append(shortNews, NewShortNews(shortNewsArr[i], tags))
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": shortNews,
		},
	)
}

func (h *Router) GetNewsCount(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filter := params.NewFilter()

	count, err := h.newsportal.GetNewsCount(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": count,
		},
	)
}
