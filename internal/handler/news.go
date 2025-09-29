package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllNews(c *gin.Context) {

	news, err := h.service.News.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": news,
		},
	)
}

func (h *Handler) GetAllNewsByQuery(c *gin.Context) {
	categoryIdStr := c.Query("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	tagIdStr := c.Query("tagId")
	tagId, _ := strconv.Atoi(tagIdStr)

	pageSizeStr := c.Query("pageSize")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)

	if categoryId == 0 || tagId == 0 ||
		pageSize == 0 || page == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "parameters not passed")
		return
	}

	news, err := h.service.News.GetAllByQuery(categoryId, tagId, pageSize, page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": news,
		},
	)
}

func (h *Handler) GetNewsById(c *gin.Context) {
	newsIdStr := c.Query("newsId")

	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	news, err := h.service.News.GetById(newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": news,
		},
	)
}

func (h *Handler) GetAllShortNews(c *gin.Context) {
	shortNews, err := h.service.News.GetAllShortNews()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": shortNews,
		},
	)
}

func (h *Handler) GetNewsCount(c *gin.Context) {
	categoryIdStr := c.Query("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	tagIdStr := c.Query("tagId")
	tagId, _ := strconv.Atoi(tagIdStr)

	count, err := h.service.News.GetCount(categoryId, tagId)
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
