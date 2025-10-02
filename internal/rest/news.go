package rest

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type queryParams struct {
	CategoryId int `form:"categoryId"`
	TagId      int `form:"tagId"`
	PageSize   int `form:"pageSize"`
	Page       int `form:"page"`
}

func (h *Handler) GetAllNews(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("parm", "CategoryId", params.CategoryId, "TagId", params.TagId, "PageSize", params.PageSize, "Page", params.Page)

	news, err := h.service.News.GetAllByQuery(params.CategoryId, params.TagId, params.PageSize, params.Page)
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
	categoryIdStr := c.Query("CategoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	tagIdStr := c.Query("TagId")
	tagId, _ := strconv.Atoi(tagIdStr)

	pageSizeStr := c.Query("PageSize")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	pageStr := c.Query("Page")
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
	newsIdStr, _ := c.Params.Get("id")

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
	categoryIdStr := c.Query("CategoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	tagIdStr := c.Query("TagId")
	tagId, _ := strconv.Atoi(tagIdStr)

	if categoryId == 0 || tagId == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "parameters were passed incorrectly")
		return
	}

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
